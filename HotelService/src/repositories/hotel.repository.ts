// Repository layer is just a layer that is used to interact with the database. It makes code more consistent,robust,extensiable, maintainable and predictable.

import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import Room from "../db/models/room";
import RoomCategory from "../db/models/roomCategory";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repository";

export class HotelRepository extends BaseRepository<Hotel> { // here we are passing Hotel in the BaseRepository class as a generic type T, so that we can use the methods of BaseRepository class on Hotel model
    // HotelRepository is a child class of BaseRepository, so it can access the methods of BaseRepository class
    // BaseRepository is a generic class, so we can pass any model to it, here
    constructor(){
        super(Hotel); // Pass the Hotel model to the BaseRepository constructor
    };

    async getAllHotel(){
        // here if you want to override the parent class method, you can do that by using the same method name in the child class
        // and if you want to any new methods in the child class, that are not in the 
        const hotels = await this.model.findAll({
            where: {
                deletedAt: null // this will fetch all the hotels which are not soft deleted
            } 
    });
    if(!hotels){
        logger.error(`Hotels not found`);
        throw new NotFoundError(`Hotels not found`);
    }
    logger.info(`Hotels found : ${hotels.length}`);
    return hotels;
 }

     async softDeleteHotelById(id: number): Promise<Hotel> {
        const hotel = await this.model.findByPk(id);
        if (!hotel) {
            logger.error(`Hotel not found : ${id}`);
            throw new NotFoundError(`Hotel with id ${id} not found`);
        }
        hotel.deletedAt = new Date();
        await hotel.save(); // Save the changes to the database
        logger.info(`Hotel soft deleted successfully, ${hotel.id}`);
        return hotel;
    }

    async getHotelById(hotelId: number) {
        console.log('fetching hotel with rooms info>>>',hotelId);
    const hotel = await this.model.findOne({
        where: { id: hotelId, deletedAt: null },
    });

    return hotel;
    }
}