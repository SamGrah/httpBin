const express = require('express');
const mongoose = require('mongoose')
const requestLogRouter = require('./controllers/requestLog')
const createBinRouter = require('./controllers/createBin');
const viewBinRouter = require('./controllers/viewBin')

const app = express();
const port = 3000;

app.use('/', express.static('../public'));

app.use(express.urlencoded({ extended: true }));
app.use(express.json());
app.use(express.raw());

mongoose.connect("mongodb://127.0.0.1:27017/Request_Records")
  .then(() => console.log('connected to MongoDB'))
  .catch(error => {
    console.log('error connecting to MongoDB', error.message)
  })

app.use('/newbin', createBinRouter)
app.use('/view-bin', viewBinRouter)  
app.use('/bin', requestLogRouter)

app.listen(port, () => console.log("RequestBin listening for requests"));