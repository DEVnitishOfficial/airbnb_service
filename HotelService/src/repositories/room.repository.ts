import { CreationAttributes } from "sequelize";
import Room from "../db/models/room";
import BaseRepository from "./base.repository";
import { Op } from "sequelize";


class RoomRepository extends BaseRepository<Room> {
    constructor() {
        super(Room);
    }

    async findRoomCategoryByIdAndDate(roomCategoryId:number, currentDate:Date){
        return await this.model.findOne({
            where : {
                roomCategoryId,
                dateOfAvailability : currentDate,
                deletedAt : null
            }
        })
    }


     async findRoomsByCategoryAndDateRange(roomCategoryId:number, startDate:Date, endDate:Date){
        return await this.model.findAll({
            where : {
                roomCategoryId,
                dateOfAvailability : {
                   [Op.between]: [startDate, endDate]
                },
                deletedAt : null
            }
        })
    }

    async bulkCreate(rooms : CreationAttributes<Room>[]){
        return await this.model.bulkCreate(rooms);
    }
}

export default RoomRepository;