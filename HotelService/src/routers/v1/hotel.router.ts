import express from "express"
import { createHotelHandler, getHotelByIdHandler } from "../../controller/hotel.controller";
import { hotelSchema } from "../../validators/hotel.validator";
import { validateRequestBody } from "../../validators";

const hotelRouter = express.Router()

hotelRouter.post('/', validateRequestBody(hotelSchema), createHotelHandler);
hotelRouter.get('/:id', getHotelByIdHandler);



export default hotelRouter;