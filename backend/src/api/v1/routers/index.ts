import {Router} from 'express'
import {logger} from '../utils/logger.util'
import * as ctrl from '../controllers/index.ctrl'

export const apiRouter = Router()
apiRouter.use(logger)
apiRouter.post('/new-bin', ctrl.createNewBin)

export const binRouter = Router() 
binRouter.use('/bin/:binId', ctrl.logRequestToBin)