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


    async findLatestDatesForAllCategories(): Promise<Array<{roomCategoryId: number, latestDate: Date}>> {
        const results = await this.model.findAll({
            where: {
                deletedAt: null
            },
            attributes: [
                'roomCategoryId',
                [this.model.sequelize!.fn('MAX', this.model.sequelize!.col('date_of_availability')), 'latestDate']
            ], // raw sql --->> SELECT roomCategoryId, MAX(date_of_availability) AS latestDate

            group: ['roomCategoryId'],
            raw: true
        });
        return results.map((result: any) => ({
            roomCategoryId: result.roomCategoryId,
            latestDate: new Date(result.latestDate)
        }));
    }

    async findByRoomCategoryIdAndDateRange(roomCategoryId:number, checkInDate:Date, checkOutDate:Date){
        return await this.model.findAll({
            where : {
                roomCategoryId,
                bookingId : null,
                dateOfAvailability : {
                   [Op.between]: [checkInDate, checkOutDate]
                }
            }
        })
    }

    async updateBookingIdToRooms(roomIds:number[], bookingId:number){
        return await this.model.update(
            { bookingId },
            {
                where: {
                    id: {
                        [Op.in]: roomIds
                    }
                }
            }
        );
    }

    async deleteBookingIdFromRooms(bookingId:number){
        return await this.model.update(
            { bookingId: null },
            {
                where: {
                    bookingId
                }
            }
        );
    }

}

export default RoomRepository;