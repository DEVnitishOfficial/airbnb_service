
import express from "express";
import { getAvailableRoomsController, updateBookingIdToRoomsController } from "../../controller/room.controller";
import { validateRequestBody } from "../../validators";
import { getAvailableRoomSchema, updateBookingIdToRoomsSchema } from "../../validators/room.validator";

const roomRouter = express.Router();

roomRouter.post("/available", validateRequestBody(getAvailableRoomSchema), getAvailableRoomsController);
roomRouter.put("/update-booking-ids", validateRequestBody(updateBookingIdToRoomsSchema), updateBookingIdToRoomsController);

export default roomRouter;
