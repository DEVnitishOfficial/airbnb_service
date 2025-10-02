
import express from 'express'
import { startRoomSchedulerController } from '../../controller/roomScheduler.controller';


const roomSchedulerRouter = express.Router();

roomSchedulerRouter.post('/start', startRoomSchedulerController)

export default roomSchedulerRouter;