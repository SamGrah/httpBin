import { Request, Response, NextFunction } from "express";
import { fetchBinRequests } from "../services/index.service";

export const getBinRequests = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  try {
    const binId = req.params.binId;
    const requests = await fetchBinRequests(binId);
    res.json(requests);
  } catch (error) {
    next(error);
  }
};
