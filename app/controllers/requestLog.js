const requestLogRouter = require('express').Router()
// const { Timestamp } = require('bson')
const requestRecords = require('../models/requestRecords')

requestLogRouter.all("/:id", async (req, res) => {
  const binExists = await requestRecords.exists({bin: req.params.id});
  if (!binExists) {
    res.end()
  }

  const recievedRequest = {
      originalUrl: req.originalUrl,
      baseUrl: req.baseUrl,
      path: req.path,
      requestType: req.method,
      headers: req.headers,
      body: req.body || undefined
  };

  await requestRecords.findOneAndUpdate(
    { bin: req.params.id },
    { $push: { requests: recievedRequest }}
  )
  res.end()
})

module.exports = requestLogRouter