import express from "express"
import { allSoftDeleteHotelHandler, createHotelHandler, getAllHotelsHandler, getHotelByIdHandler, hardDeleteHotelHandler, restoreSoftDeletedHotelByIdHandler, softDeleteHotelHandler, updateHotelHandler } from "../../controller/hotel.controller";
import { hotelSchema } from "../../validators/hotel.validator";
import { validateRequestBody } from "../../validators";

const hotelRouter = express.Router()

hotelRouter.post('/', validateRequestBody(hotelSchema), createHotelHandler);
hotelRouter.get('/allhotel', getAllHotelsHandler);
hotelRouter.get('/allSoftDeletedHotels', allSoftDeleteHotelHandler);
hotelRouter.get('/id/:id', getHotelByIdHandler);
hotelRouter.put('/updateById/:id', updateHotelHandler);
hotelRouter.delete('/softDeleteById/:id',softDeleteHotelHandler);
hotelRouter.delete('/hardDeleteById/:id', hardDeleteHotelHandler);
hotelRouter.put('/restoreSoftDeletedHotelById/:id', restoreSoftDeletedHotelByIdHandler);


export default hotelRouter;