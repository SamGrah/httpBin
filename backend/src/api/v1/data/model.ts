import mongoose, { Schema } from "mongoose";

const requestRecordsSchema: Schema = new mongoose.Schema(
  {
    bin: String,
    requests: [
      {
        originalUrl: String,
        baseUrl: String,
        path: String,
        requestType: String,
        headers: { type: {}, default: {} },
        body: { type: {}, default: {} },
        senderIP: String,
      },
    ],
    expireAt: {
      type: Date,
      default: Date.now,
      index: { expires: 60 * 60 * 24 * 2 },
    },
  },
  { minimize: false }
);

export default requestRecordsSchema;
