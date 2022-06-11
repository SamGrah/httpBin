import { Response, NextFunction } from "express";
import { createUniqueBin } from "../services/create_unique_bin.service";

export const createNewBin = async (_, res: Response, next: NextFunction) => {
  try {
    const newBin = await createUniqueBin();
    res.json(newBin);
  } catch (error) {
    next(error);
  }
};
