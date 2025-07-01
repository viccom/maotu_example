const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 3001 });

let boolVar = false;

wss.on('connection', (ws) => {
  console.log('WebSocket client connected');
  // 新增：连接建立时立即推送一次当前数据
  const data = {
    floatVar: randomFloat().toFixed(2), // 保留两位小数
    boolVar
  };
  ws.send(JSON.stringify(data));
  ws.on('close', () => {
    console.log('WebSocket client disconnected');
  });
  ws.on('error', (err) => {
    console.error('WebSocket client error:', err);
  });
});

setInterval(() => {
  boolVar = !boolVar;
}, 60000);

function randomFloat() {
  return Math.random() * 100;
}

setInterval(() => {
  const data = {
    floatVar: randomFloat().toFixed(2), // 保留两位小数
    boolVar
  };
  const msg = JSON.stringify(data);
  wss.clients.forEach(client => {
    if (client.readyState === WebSocket.OPEN) {
      client.send(msg);
    }
  });
  console.log('Sent data to clients:', msg);
}, 1000);

console.log('WebSocket server running at ws://localhost:3001');
