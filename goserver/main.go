package main

import (
	"embed"
	"encoding/json"
	"flag"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

	//go:embed dist
	distFS embed.FS
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

// --- 新增：静态文件服务中间件 ---
func staticFileServer() gin.HandlerFunc {
	// 1. 从 embed.FS 创建一个子文件系统，并创建一个文件服务器
	staticFS, err := fs.Sub(distFS, "dist")
	if err != nil {
		log.Fatal("Failed to create sub FS for dist:", err)
	}
	fileServer := http.FileServer(http.FS(staticFS))

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// 2. 检查请求是否是 API 或 WebSocket，如果是，则跳过中间件
		if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/ws") {
			c.Next()
			return
		}

		// 3. 检查文件是否存在于我们的嵌入式文件系统中
		// 我们需要去掉路径开头的'/'，因为fs.FS不使用它
		filePath := strings.TrimPrefix(path, "/")
		// 如果是根路径，查找 index.html
		if filePath == "" {
			filePath = "index.html"
		}

		f, err := staticFS.Open(filePath)
		if err != nil {
			// 如果文件不存在，我们假设它是一个 SPA 路由
			// 把 index.html 的内容返回
			data, err := distFS.ReadFile("dist/index.html")
			if err != nil {
				log.Printf("Could not read index.html: %v", err)
				c.String(http.StatusInternalServerError, "index.html not found")
				return
			}
			c.Data(http.StatusOK, "text/html; charset=utf-8", data)
			c.Abort() // 终止请求链
			return
		}
		f.Close() // 检查完就关闭

		// 4. 如果文件确实存在，让 http.FileServer 来处理它
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort() // 终止请求链，防止传递给 Gin 的 NoRoute
	}
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}
	_ = err
}

func main() {
	// 解析命令行参数
	addr := flag.String("addr", ":3000", "监听地址")
	autoOpen := flag.Bool("open", true, "启动后自动打开浏览器")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	loadDevices()
	simDataMap = make(map[string]interface{})
	startSimulate()

	router := gin.Default()

	// --- 核心改动：使用我们的自定义中间件 ---
	// 这个中间件会处理所有非 API 的 GET 请求
	router.Use(staticFileServer())

	// 现在可以安全地注册 API 路由，不会有任何冲突
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/devices", func(c *gin.Context) {
			c.JSON(http.StatusOK, devicesList)
		})
		apiGroup.GET("/datas", func(c *gin.Context) {
			simDataMutex.RLock()
			defer simDataMutex.RUnlock()
			c.JSON(http.StatusOK, simDataMap)
		})
		apiGroup.POST("/saveview", func(c *gin.Context) {
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
		apiGroup.GET("/loadview", func(c *gin.Context) {
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
	}

	// 注册 gin WebSocket 路由
	router.GET("/ws", func(c *gin.Context) {
		wsHandler(c)
	})

	// 因为我们的中间件已经处理了所有情况，所以 NoRoute 变得不再必要
	// 你可以移除它，或者保留它来捕获非 GET 的未定义路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "ROUTE_NOT_FOUND", "message": "Endpoint not found"})
	})

	// 启动 HTTP 服务
	go func() {
		log.Println("请访问 http://localhost:3000/api/devices 获取设备列表")
		log.Println("请访问 http://localhost:3000/api/datas 获取仿真数据")
		log.Println("HTTP server running at http://localhost" + *addr)
		log.Println("Access the web UI at http://localhost" + *addr)
		log.Println("Press Ctrl+C to stop the server")
		if err := router.Run(*addr); err != nil {
			log.Fatal(err)
		}
	}()

	if *autoOpen {
		go openBrowser("http://127.0.0.1" + *addr)
	}
	// 启动独立的 WebSocket 服务（这个部分保持不变）
	http.HandleFunc("/", wsHandlerRaw)
	log.Println("请连接 ws://localhost:3001 获取实时仿真数据")
	log.Println("WebSocket server running at ws://localhost:3001")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
