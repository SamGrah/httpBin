import { Request, Response, NextFunction } from "express";
import { logRequest } from "../services/log_request.service";

export const logRequestToBin = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  try {
    // let body = req.body;
    // if (req.body.payload) body = JSON.parse(req.body.payload);

    // const requestProps = {
    //   originalUrl: req.originalUrl,
    //   baseUrl: req.baseUrl,
    //   path: req.path,
    //   requestType: req.method,
    //   headers: req.headers,
    //   body: body,
    //   senderIP: req.headers['x-forwarded-for']
    // };

    // const binId = req.params.binId
    // const newBin = await logRequest(binId, requestProps);
    // res.json(newBin);
    
  } catch (error) {
    next(error);
  }
};
