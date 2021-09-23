var WebSocketServer = require('ws').Server
wss = new WebSocketServer({port: 40510});
const websockets = {};

wss.on('connection', function (ws) {
  let bin
  ws.on('message', function (binId) {
    bin = binId
    websockets[bin] = websockets[bin] || [];
    websockets[bin].push(ws);
  });

  ws.on('close', function(binId) {
    let idx = websockets[bin].findIndex(socketObj => {
      if (socketObj == ws) return true
    })
    
    websockets[bin].splice(idx, 1)
    if (websockets[bin].length == 0) delete websockets[bin]
  })
})

module.exports = websockets  