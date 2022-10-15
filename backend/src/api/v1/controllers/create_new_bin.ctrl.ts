import { Response, NextFunction } from "express";
import { IRequest } from "../routers";
import { createUniqueBin } from "../services/create_unique_bin.service";

export const createNewBin = async (_: IRequest, res: Response, next: NextFunction) => {
  try {
    // const newBin = await createUniqueBin();
    // res.json(newBin);
    res.json(_.test)
  } catch (error) {
    next(error);
  }
};
