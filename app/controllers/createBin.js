const createBinRouter = require('express').Router()
const requestRecords = require('../models/requestRecords')

createBinRouter.get('/', async (req, res) => {
  const generateRandomBin = () => {
    result = ''
    const chars = 'abcdefghijklmnopqrstuvwxyx0123456789'
    for ( let idx = 0; idx < 6; idx += 1) {
      result += chars.charAt(Math.floor(Math.random() * chars.length))
    }
    return result
  };
  
  req.hostname

  let binId = generateRandomBin()
  while (await requestRecords.exists({bin: binId})) 
    binId = generateRandomBin()

  await requestRecords.create({bin: binId, requests: []})
  res.render('new-bin',{
    binId: binId,
    hostname: req.hostname, 
  });
});


module.exports = createBinRouter
