const viewBinRouter = require('express').Router()
const requestRecords = require('../models/requestRecords')

viewBinRouter.get('/:binId', async (req, res) => {
  const binExists = await requestRecords.exists({bin: req.params.binId});
  if (!binExists) {
    res.redirect('/')
    res.end()
  } else {
    const binContents = await requestRecords.find({bin: req.params.binId}).lean()
    res.render('view-bin', {
      "bin": binContents[0].bin,
      "requests": binContents[0].requests
    })
  }
})


module.exports = viewBinRouter