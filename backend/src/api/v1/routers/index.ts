import {Router} from 'express'
import {logger} from '../utils/logger.util'
import * as ctrl from '../controllers/index.ctrl'

export const router = Router()

router.use(logger)
router.post('/new-bin', ctrl.createNewBin)