// Repository layer is just a layer that is used to interact with the database. It makes code more consistent,robust,extensiable, maintainable and predictable.

import { Op } from "sequelize";
import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { createHotelDto, updateHotelDto } from "../dto/hotel.dto";
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
    const hotel = await Hotel.findOne({
        where: {
            id: id,
            deletedAt: null // this will fetch the hotel which is not soft deleted
        }
    });
    if(!hotel){
        logger.error(`Hotel not found : ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`)
    }
    return hotel;
}
export async function getAllHotel(){
    const hotels = await Hotel.findAll({
        where: {
            deletedAt: null // this will fetch all the hotels which are not soft deleted
        }
    });
    logger.info(`Fetching all hotels completed`,hotels);
    if(!hotels){
        logger.error(`Hotels not found`);
        throw new NotFoundError(`Hotels not found`)
    }
    return hotels;
}

export async function updateHotelById(id:number, hotelData:updateHotelDto){
    const hotel = await Hotel.findOne({
        where: {
            id: id,
            deletedAt: null // this will fetch the hotel which is not soft deleted ;
        }
    });
    if(!hotel){
        logger.error(`Hotel not found : ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`)
    }
    const updatedHotel = await hotel.update(hotelData);
    return updatedHotel;
}

export async function softDeleteHotelById(id:number){
    const hotel = await Hotel.findByPk(id);
    if(!hotel){
        logger.error(`Hotel not found : ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`)
    }
    hotel.deletedAt = new Date();
    await hotel.save(); // Save the changes to the database
    logger.info(`Hotel soft deleted successfully, ${hotel.id}`);
    return hotel;
}

export async function hardDeleteHotelById(id:number){
    const hotel = await Hotel.findByPk(id);
    if(!hotel){
        logger.error(`Hotel not found : ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`)
    }
    await hotel.destroy();
    logger.info(`Hotel hard deleted successfully, ${hotel.id}`);
    return hotel;
}

export async function allSoftDeletedHotels(){
    const hotels = await Hotel.findAll({
        where: {
            deletedAt: {
                // Op --> sequelize operator
                // Op.ne stands for "not equal" (ne = not equal).
                // overall it says to sequelize "Give me all rows from the Hotel table where deletedAt is not null."
                [Op.ne]: null // this will fetch all the hotels which are soft deleted
            }
        }
    });
    logger.info(`Fetching all deleted hotels completed`,hotels);
    if(!hotels){
        logger.error(`Deleted Hotels not found`);
        throw new NotFoundError(`Deleted Hotels not found`)
    }
    return hotels;
}

export async function restoreSoftDeletedHotelById(id:number){
    const hotel = await Hotel.findOne({
        where:{
            id: id,
            deletedAt: {
                [Op.ne]: null // this will fetch the hotel which is soft deleted
            }
        }
    })
    if(!hotel){
        logger.error(`Hotel not found : ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`)
    }
    hotel.deletedAt = null; // set deletedAt to null to restore the hotel
    await hotel.save(); // Save the changes to the database
    logger.info(`Hotel restored successfully, ${hotel.id}`);
    return hotel;
}
