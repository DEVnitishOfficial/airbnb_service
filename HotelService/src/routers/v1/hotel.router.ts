import express from "express"
import { createHotelHandler, deleteHotelHandler, getAllHotelsHandler, getHotelByIdHandler, updateHotelHandler } from "../../controller/hotel.controller";
import { hotelSchema } from "../../validators/hotel.validator";
import { validateRequestBody } from "../../validators";

const hotelRouter = express.Router()

hotelRouter.post('/', validateRequestBody(hotelSchema), createHotelHandler);
hotelRouter.get('/allhotel', getAllHotelsHandler);
hotelRouter.get('/id/:id', getHotelByIdHandler);
hotelRouter.put('/updateById/:id', updateHotelHandler);
hotelRouter.delete('/deleteById/:id', deleteHotelHandler);



export default hotelRouter;