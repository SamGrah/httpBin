// const { Timestamp } = require('bson')
const mongoose = require('mongoose')

const requestRecordsSchema = new mongoose.Schema({
  bin : String,
  requests: {
    originalUrl: String,
    baseUrl: String,
    path: String,
    requestType: String,
    headers: {},
    body: {},
    encoding: String
  }
})

module.exports = mongoose.model('Requests', requestRecordsSchema)