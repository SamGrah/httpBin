import { IncomingHttpHeaders } from "http";
import mongoose from "mongoose";

export type RequestProperties = {
  originalUrl: string;
  baseUrl: string;
  path: string;
  requestType: string;
  headers: IncomingHttpHeaders;
  body: { [key: string]: string };
  senderIP: string | string[];
};

export type HttpRequest = mongoose.Document & {
  bin: string;
  requests: RequestProperties[];
};

const requestRecordsSchema = new mongoose.Schema(
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

export const requestRecords = mongoose.model<HttpRequest>(
  "httpRequest",
  requestRecordsSchema
);
