// Repository layer is just a layer that is used to interact with the database. It makes code more consistent,robust,extensiable, maintainable and predictable.

import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { createHotelDto } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";

// here we have to provide the datatype of the hotelData which will come from the client side either from the browser or the postman, so for the transferable data we create dto(data transfer object) layer hence created dto folder and inside that written the types of hoteldata
export async function createHotel(hotelData : createHotelDto) {
    const hotel = await Hotel.create({
        name : hotelData.name,
        address: hotelData.address,
        location: hotelData.location,
        rating: hotelData.rating,
        ratingCount: hotelData.ratingCount 
    });
    logger.info(`hotel created successfully, ${hotel.id}`);
    return hotel;
}

export async function getHotelById(id:number){
    const hotel = await Hotel.findByPk(id);
    if(!hotel){
        logger.error(`Hotel not found : ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`)
    }
    return hotel;
}
