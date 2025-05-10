import express from "express"
import { createHotelHandler, deleteHotelHandler, getAllHotelsHandler, getHotelByIdHandler, updateHotelHandler } from "../../controller/hotel.controller";
import { hotelSchema } from "../../validators/hotel.validator";
import { validateRequestBody } from "../../validators";

const hotelRouter = express.Router()

hotelRouter.post('/', validateRequestBody(hotelSchema), createHotelHandler);
hotelRouter.get('/allhotel', getAllHotelsHandler);
hotelRouter.put('/updateById/:id', updateHotelHandler);
hotelRouter.delete('/deleteById/:id', deleteHotelHandler);
hotelRouter.get('/id/:id', getHotelByIdHandler);


export default hotelRouter;