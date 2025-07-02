package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type DeviceChild struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Dtype string `json:"dtype"`
}
type Device struct {
	Label    string        `json:"label"`
	Value    string        `json:"value"`
	Children []DeviceChild `json:"children"`
}

var (
	devicesList  []Device
	simDataMap   map[string]interface{}
	simDataMutex sync.RWMutex
	viewJsonPath string
	devicesPath  string
	upgrader     = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	wsClients     = make(map[*websocket.Conn]bool)
	wsClientsLock sync.Mutex
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var baseDir string
	// go run . 时，始终用当前工作目录
	wd, _ := os.Getwd()
	baseDir = wd
	log.Println("配置文件查找基准目录:", baseDir)

	// 优先查找 ./config/devices.json，否则查找 ./devices.json
	devicesPath = filepath.Join(baseDir, "config", "devices.json")
	viewJsonPath = filepath.Join(baseDir, "config", "view.json")
	if _, err := os.Stat(devicesPath); os.IsNotExist(err) {
		devicesPath = filepath.Join(baseDir, "devices.json")
	}
	if _, err := os.Stat(viewJsonPath); os.IsNotExist(err) {
		viewJsonPath = filepath.Join(baseDir, "view.json")
	}
}

// 判断是否 go run . 运行
func isGoRunMode(exePath string) bool {
	return filepath.Base(exePath) == "___go_build_main_go.exe" || // Windows
		filepath.Base(exePath) == "___go_build_main_go" || // Linux/Mac
		filepath.Base(filepath.Dir(exePath)) == "go-build" // 兼容 go1.21+
}

func loadDevices() {
	data, err := ioutil.ReadFile(devicesPath)
	if err != nil {
		log.Println("加载 devices.json 失败:", err)
		devicesList = []Device{}
		return
	}
	if err := json.Unmarshal(data, &devicesList); err != nil {
		log.Println("解析 devices.json 失败:", err)
		devicesList = []Device{}
	}
}

func randomFloat() float64 {
	return float64(int(rand.Float64()*10000)) / 100
}
func randomInt() int {
	return rand.Intn(101)
}
func randomBool() bool {
	return rand.Intn(2) == 0
}
func simulateData() map[string]interface{} {
	result := make(map[string]interface{})
	for _, device := range devicesList {
		for _, child := range device.Children {
			switch child.Dtype {
			case "float":
				result[child.Value] = randomFloat()
			case "int":
				result[child.Value] = randomInt()
			case "bool":
				result[child.Value] = randomBool()
			default:
				result[child.Value] = nil
			}
		}
	}
	return result
}

func startSimulate() {
	go func() {
		for {
			simDataMutex.Lock()
			simDataMap = simulateData()
			simDataMutex.Unlock()
			b, _ := json.Marshal(simDataMap)
			broadcastWS(b)
			time.Sleep(time.Second)
		}
	}()
}

func broadcastWS(msg []byte) {
	wsClientsLock.Lock()
	defer wsClientsLock.Unlock()
	for c := range wsClients {
		err := c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			c.Close()
			delete(wsClients, c)
		}
	}
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	wsClientsLock.Lock()
	wsClients[conn] = true
	wsClientsLock.Unlock()
	// 新连接立即推送一次仿真数据
	simDataMutex.RLock()
	initData, _ := json.Marshal(simDataMap)
	simDataMutex.RUnlock()
	conn.WriteMessage(websocket.TextMessage, initData)
	for {
		// 只读，不处理客户端消息
		if _, _, err := conn.NextReader(); err != nil {
			break
		}
	}
	wsClientsLock.Lock()
	delete(wsClients, conn)
	wsClientsLock.Unlock()
	conn.Close()
}

// 新增：原生 WebSocket 处理函数
func wsHandlerRaw(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	wsClientsLock.Lock()
	wsClients[conn] = true
	wsClientsLock.Unlock()
	// 新连接立即推送一次仿真数据
	simDataMutex.RLock()
	initData, _ := json.Marshal(simDataMap)
	simDataMutex.RUnlock()
	conn.WriteMessage(websocket.TextMessage, initData)
	for {
		if _, _, err := conn.NextReader(); err != nil {
			break
		}
	}
	wsClientsLock.Lock()
	delete(wsClients, conn)
	wsClientsLock.Unlock()
	conn.Close()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	loadDevices()
	simDataMap = make(map[string]interface{})
	startSimulate()

	router := gin.Default()

	// HTTP API
	router.GET("/api/devices", func(c *gin.Context) {
		c.JSON(http.StatusOK, devicesList)
	})
	router.GET("/api/datas", func(c *gin.Context) {
		simDataMutex.RLock()
		defer simDataMutex.RUnlock()
		c.JSON(http.StatusOK, simDataMap)
	})
	router.POST("/api/saveview", func(c *gin.Context) {
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "参数错误"})
			return
		}
		b, _ := json.MarshalIndent(body, "", "  ")
		if err := ioutil.WriteFile(viewJsonPath, b, 0644); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "保存失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "配置已保存"})
	})
	router.GET("/api/loadview", func(c *gin.Context) {
		if _, err := os.Stat(viewJsonPath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "未找到 view.json"})
			return
		}
		b, err := ioutil.ReadFile(viewJsonPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "读取失败"})
			return
		}
		var viewConfig interface{}
		if err := json.Unmarshal(b, &viewConfig); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "解析失败"})
			return
		}
		c.JSON(http.StatusOK, viewConfig)
	})

	// WebSocket
	router.GET("/ws", func(c *gin.Context) {
		wsHandler(c)
	})

	// 启动 HTTP 服务
	go func() {
		log.Println("HTTP API server running at http://localhost:3000")
		log.Println("请访问 http://localhost:3000/api/devices 获取设备列表")
		log.Println("请访问 http://localhost:3000/api/datas 获取仿真数据")
		if err := router.Run(":3000"); err != nil {
			log.Fatal(err)
		}
	}()

	// 启动 WebSocket 服务（ws://localhost:3001）
	http.HandleFunc("/", wsHandlerRaw)
	log.Println("WebSocket server running at ws://localhost:3001")
	log.Println("请连接 ws://localhost:3001 获取实时仿真数据")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
