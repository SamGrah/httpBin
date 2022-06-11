import { Request, NextFunction } from "express";

export const logger = (req: Request, _, next: NextFunction) => {
  console.log(`${req.method} ${req.url} ${req.path}`);
  next();
};
