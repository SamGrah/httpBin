import {Response} from 'express';
import { hostname } from 'os';

export const getHostname = (_, res: Response) => {
  res.json(hostname())
}