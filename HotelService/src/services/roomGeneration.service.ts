import { CreationAttributes } from "sequelize";
import RoomCategory from "../db/models/roomCategory";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import RoomRepository from "../repositories/room.repository";
import RoomCategoryRepository from "../repositories/roomCategory.repository";
import { BadRequesError, NotFoundError } from "../utils/errors/app.error";
import Room from "../db/models/room";
import logger from "../config/logger.config";



const roomCategoryRepository = new RoomCategoryRepository()
const roomRepository = new RoomRepository()

export async function generateRoomsService(jobData: RoomGenerationJob) {

    console.log('request recieved at service', jobData);

    let totalRoomsCreated = 0;
    let totalDatesProcessed = 0;

    // check if the category exists

    const roomCategory = await roomCategoryRepository.findById(jobData.roomCategoryId);

    if (!roomCategory) {
        throw new NotFoundError(`Room category with id ${jobData.roomCategoryId} not found`)
    }

    const startDate = new Date(jobData.startDate)
    const endDate = new Date(jobData.endDate)


    if (startDate >= endDate) {
        throw new BadRequesError("Start date must be before end date")
    }

    if (startDate < new Date()) {
        throw new BadRequesError("Start date must be in the future")
    }

    // total days for which we have to create rooms


    const totalDays = Math.floor((endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24));

    logger.info(`Generating rooms for ${totalDays} days`)

    const batchSize = jobData.batchSize || 100 // put this one in the config or .env

    const currentDate = new Date(startDate)

    while(currentDate <= endDate){
        const batchEndDate = new Date(currentDate);

        batchEndDate.setDate(batchEndDate.getDate() + batchSize - 1); // -1 because currentDate is inclusive

        if(batchEndDate > endDate){ // suppose we have to create rooms for 130 days and batch size is 100, then in second batch we have to create rooms for 30 days only but the above line batchEndDate.setDate..... will a batch of 200, this things i want to avoid that's why it's necessery to check endDate.
            batchEndDate.setTime(endDate.getTime());
        }


        const batchResult = await processDateBatch(roomCategory, currentDate, batchEndDate, jobData.priceOverride)

        console.log('see the batch result', batchResult);

        totalRoomsCreated += batchResult.roomsCreated;
        totalDatesProcessed += batchResult.dateProcessed;
        
        currentDate.setTime(batchEndDate.getTime() + 1); // move to the next day after batchEndDate

        // From the generateRoomsService we are encountring infinite loop because we are setting currentDate = batchEndDate and sometimes batchEndDate === endDate.

        // This means currentDate <= endDate never becomes false, so the loop repeats infinitely.

        // Solution: always advance currentDate to the next day after the batch: that's why added + 1 and for including current date used -1 on line No 56
    }

    return{
        totalRoomsCreated,
        totalDatesProcessed
    }


}

export async function processDateBatch(roomCategory: RoomCategory, startDate: Date, endDate: Date, priceOverride?: number) {
    let roomsCreated = 0;
    let dateProcessed = 0;
    const roomsToCreate: CreationAttributes<Room>[] = [];

    const currentDate = new Date(startDate)

    // Below code is the problem of n+1 query
    // TODO : use a better query to get rooms
    // here we are making seperate db query for each day, if we have to create next 100 rooms
    // then here we are making 100 db query just for cheking if the rooms are already existing on that date or not.

    // while(currentDate <= endDate){
    //     const existingRoom = await roomRepository.findRoomCategoryByIdAndDate(roomCategory.id, currentDate);

    //     if(!existingRoom){
    //         roomsToCreate.push({
    //             hotelId : roomCategory.hotelId,
    //             roomCategoryId : roomCategory.id,
    //             dateOfAvailability : currentDate,
    //             price : priceOverride || roomCategory.price
    //         })
    //     }

    //     currentDate.setDate(currentDate.getDate() + 1);
    //     dateProcessed;
    // }




    /** Better query solution */

    // How the below solution solve n+1 queries problem
    //Instead of calling findRoomCategoryByIdAndDate for each day separately, fetch all existing rooms for the date range in one query:
    
    const existingRooms = await roomRepository.findRoomsByCategoryAndDateRange(
        roomCategory.id,
        startDate,
        endDate
    );

    // put them in a Set for quick lookup
    const existingDates = new Set(
        existingRooms.map(r => new Date(r.dateOfAvailability).toISOString().split('T')[0])
    );

    while (currentDate <= endDate) {
        const dateKey = currentDate.toISOString().split('T')[0];
        if (!existingDates.has(dateKey)) {
            roomsToCreate.push({
                hotelId: roomCategory.hotelId,
                roomCategoryId: roomCategory.id,
                dateOfAvailability: new Date(currentDate),
                price: priceOverride || roomCategory.price,
                roomNo: 1 + roomsCreated, // Assigning room number sequentially
                createdAt : new Date(),
                updatedAt : new Date(),
                deletedAt : null
            });
        }
        currentDate.setDate(currentDate.getDate() + 1);
        dateProcessed++;
    }

    if (roomsToCreate.length > 0) {
         logger.info(`Creating ${roomsToCreate.length} rooms`);
        await roomRepository.bulkCreate(roomsToCreate)
        roomsCreated += roomsToCreate.length;
    }

    return{
        roomsCreated,
        dateProcessed
    }

}