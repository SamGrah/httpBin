const requestLogRouter = require('express').Router()
const requestRecords = require('../models/requestRecords')

requestLogRouter.all("/:id", async (req, res) => {
  const binExists = await requestRecords.exists({bin: req.params.id});
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
    { bin: req.params.id },
    { $push: 
      { 
        requests: {
          $each: [ recievedRequest],
          $position: 0 
        }
      }
    }
  )
  res.end()
})

module.exports = requestLogRouter