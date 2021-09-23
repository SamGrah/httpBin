const requestLogRouter = require('express').Router()
const requestRecords = require('../models/requestRecords')
const openWebSockets = require('../websocket/websocket')

requestLogRouter.all("/:id", async (req, res) => {
  const binId = req.params.id
  const binExists = await requestRecords.exists({bin: binId});
  if (!binExists) {
    res.end()
  }

  let body = req.body
  if (req.body.payload) body = JSON.parse(req.body.payload)
  
  const recievedRequest = {
      originalUrl: req.originalUrl,
      baseUrl: req.baseUrl,
      path: req.path,
      requestType: req.method,
      headers: req.headers,
      body: body,
      senderIP: req.headers['x-forwarded-for']
  };

  await requestRecords.findOneAndUpdate(
    { bin: binId },
    { $push: 
      { 
        requests: {
          $each: [ recievedRequest],
          $position: 0,
          $slice: 20 
        }
      }
    }
  )
  
  if (openWebSockets[binId]) {
    openWebSockets[binId].forEach(webSocket => {
      webSocket.send('refresh')
    })
  }
  res.end()
})

module.exports = requestLogRouter