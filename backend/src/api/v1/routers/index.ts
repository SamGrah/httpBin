import { Router, Request } from "express";
import { logger } from "../utils/logger.util";
import * as ctrl from "../controllers/index.ctrl";

export interface IRequest extends Request {
  test: string
}
export const apiRouter = Router();
apiRouter.use(logger);
apiRouter.use((req: IRequest, _, next) => {
  req.params["test"] = "test";
  ((req, _, next) => {
    req.test = 'test'
    next()
  })(req, _, next);
});

apiRouter.post("/new-bin", ctrl.createNewBin);

export const binRouter = Router();
binRouter.use("/bin/:binId", ctrl.logRequestToBin);
