import express, { Application, Request, Response } from "express";
import {router} from './routers/index'
import mongoose from "mongoose";
import cors from 'cors';

const dotenv = require("dotenv");
dotenv.config();

const app = express();

app.use(cors())
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use('/api/test', (_, res) => {
  res.json('httpBin Api Is Running')
})
app.use('/api', router)

mongoose
  .connect(process.env.DATABASE ?? "")
  .then(async () => {
    // await dbInit() 
    console.log("connected to MongoDB")
  })
  .catch((error) => {
    console.log("error connecting to MongoDB", error.message);
  });

export default app;
