// const { Timestamp } = require('bson')
const mongoose = require('mongoose')

const requestRecordsSchema = new mongoose.Schema({
  bin : String,
  requests: [{
    originalUrl: String,
    baseUrl: String,
    path: String,
    requestType: String,
    headers: { type: {}, default: {} },
    body: { type: {}, default: {} },
    senderIP: String
  }]
}, { minimize: false})

module.exports = mongoose.model('Requests', requestRecordsSchema)