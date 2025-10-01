import express from 'express'
import { validateRequestBody } from '../../validators';
import { RoomGenerationJobSchema } from '../../dto/roomGeneration.dto';
import { generateRoomsController } from '../../controller/roomGeneration.controller';

const roomGenerationRouter = express.Router();

roomGenerationRouter.post('/', validateRequestBody(RoomGenerationJobSchema), generateRoomsController);

export default roomGenerationRouter;