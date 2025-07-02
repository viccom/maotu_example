const fs = require('fs');
const path = require('path');
const express = require('express');
const WebSocket = require('ws');

// 1. 加载 devices.json
const devicesPath = path.join(__dirname, 'devices.json');
let devicesList = [];
try {
  const raw = fs.readFileSync(devicesPath, 'utf-8');
  devicesList = JSON.parse(raw);
} catch (e) {
  console.error('加载 devices.json 失败:', e);
  devicesList = [];
}

// 2. 仿真数据 map
let simDataMap = {};

// 3. 启动 http api 服务
const app = express();
const HTTP_PORT = 3000;
app.get('/api/devices', (req, res) => {
  res.json(devicesList);
});
app.get('/api/datas', (req, res) => {
  res.json(simDataMap);
});
app.listen(HTTP_PORT, () => {
  console.log(`HTTP API server running at http://localhost:${HTTP_PORT}`);
  console.log('请访问 http://localhost:3000/api/devices 获取设备列表');
  console.log('请访问 http://localhost:3000/api/datas 获取仿真数据');
});

// 4. 启动 WebSocket 服务
const WS_PORT = 3001;
const wss = new WebSocket.Server({ port: WS_PORT });
wss.on('connection', (ws) => {
  console.log('WebSocket client connected');
  // 新连接立即推送一次仿真数据
  ws.send(JSON.stringify(simDataMap));
  ws.on('close', () => {
    console.log('WebSocket client disconnected');
  });
  ws.on('error', (err) => {
    console.error('WebSocket client error:', err);
  });
});
console.log(`WebSocket server running at ws://localhost:${WS_PORT}`);
console.log('请连接 ws://localhost:3001 获取实时仿真数据');

// 5. 仿真线程：每秒生成一次仿真数据
function randomFloat() {
  return +(Math.random() * 100).toFixed(2);
}
function randomInt() {
  return Math.floor(Math.random() * 101);
}
function randomBool() {
  return Math.random() < 0.5;
}
function simulateData() {
  const map = {};
  devicesList.forEach(device => {
    if (Array.isArray(device.children)) {
      device.children.forEach(child => {
        let val;
        switch (child.dtype) {
          case 'float':
            val = randomFloat();
            break;
          case 'int':
            val = randomInt();
            break;
          case 'bool':
            val = randomBool();
            break;
          default:
            val = null;
        }
        map[child.value] = val;
      });
    }
  });
  return map;
}
setInterval(() => {
  simDataMap = simulateData();
  // 通过 websocket 推送数据
  const msg = JSON.stringify(simDataMap);
  wss.clients.forEach(client => {
    if (client.readyState === WebSocket.OPEN) {
      client.send(msg);
    }
  });
}, 1000);
