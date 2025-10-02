import  cron,  {ScheduledTask} from 'node-cron'
import logger from '../config/logger.config';
import RoomRepository from '../repositories/room.repository';
import RoomCategoryRepository from '../repositories/roomCategory.repository';
import { RoomGenerationJob } from '../dto/roomGeneration.dto';
import { addRoomGenerationJobToQueue } from '../producers/roomGeneration.producer';
import { serverConfig } from '../config';


const roomRepository = new RoomRepository();
const roomCategoryRepository = new RoomCategoryRepository();


let cronJob : ScheduledTask | null = null

export async function startRoomSchedulerCronJob() : Promise<void> {


    if (cronJob) {
        logger.warn('Room scheduler is already running');
        return;
    }

     // Schedule job to run every minute
    cronJob = cron.schedule(serverConfig.ROOM_CRON, async () => {
        logger.info('Starting room availability extension check');
        await extendRoomAvailability()
        logger.info('Room availability extension check completed');
    },{timezone: 'UTC'});

    cronJob.start();
    logger.info(`Room availability extension scheduler started - running every ${"* * * * *"}`);
}


export async function extendRoomAvailability() : Promise<void> {


   const roomCategoriesWithLatestDates = await roomRepository.findLatestDatesForAllCategories();
    

        if (roomCategoriesWithLatestDates.length === 0) {
            logger.info('No room categories found with availability dates');
            return;
        }

        logger.info(`Found ${roomCategoriesWithLatestDates.length} room categories to extend`);

        // Process each room category
        for (const categoryData of roomCategoriesWithLatestDates) {
            await extendCategoryAvailability(categoryData);
        }
}


export async function extendCategoryAvailability(categoryData: { roomCategoryId: number, latestDate: Date }): Promise<void>{


    try {
        const { roomCategoryId, latestDate } = categoryData;

        // Calculate the next date (one day after the latest date)
        const nextDate = new Date(latestDate);
        nextDate.setDate(nextDate.getDate() + 1);

        // Check if the room category still exists
        const roomCategory = await roomCategoryRepository.findById(roomCategoryId);


        if (!roomCategory) {
            logger.warn(`Room category ${roomCategoryId} not found, skipping extension`);
            return;
        }

        // Check if room for next date already exists
        const existingRoom = await roomRepository.findRoomCategoryByIdAndDate(roomCategoryId, nextDate);
        if (existingRoom) {
            logger.debug(`Room for category ${roomCategoryId} on ${nextDate.toISOString()} already exists, skipping`);
            return;
        }

        const endDate = new Date(nextDate);
        endDate.setDate(endDate.getDate() + 1);

        // Create job to generate room for the next date
        const jobData: RoomGenerationJob = {
            roomCategoryId: roomCategoryId,
            startDate: nextDate.toISOString(),
            endDate: endDate.toISOString(),
            priceOverride: roomCategory.price,
            batchSize: 1
        };

        // Add job to queue
        await addRoomGenerationJobToQueue(jobData);
        
        logger.info(`Added room generation job for category ${roomCategoryId} on ${nextDate.toISOString()}`);

    } catch (error) {
        logger.error(`Error extending availability for room category ${categoryData.roomCategoryId}:`, error);
        // Don't throw here to avoid stopping the entire scheduler
    }
};