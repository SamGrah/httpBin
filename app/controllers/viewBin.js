const viewBinRouter = require('express').Router()
// const { Timestamp } = require('bson')
const requestRecords = require('../models/requestRecords')

viewBinRouter.get('/:binId', async (req, res) => {
  const binExists = await requestRecords.exists({bin: req.params.binId});
  if (!binExists) {
    res.redirect('/')
    res.end()
  } else {
    const binContents = await requestRecords.find({bin: req.params.binId}) 
    res.send(binContents)
  }
})


module.exports = viewBinRouter