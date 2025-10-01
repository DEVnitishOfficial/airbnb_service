import { Job, Worker } from "bullmq";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import { ROOM_GENERATION_QUEUE } from "../queues/roomGeneration.queue";
import { getRedisConnObject } from "../config/redis.config";
import { ROOM_GENERATION_PAYLOAD } from "../producers/roomGeneration.producer";
import { generateRoomsService } from "../services/roomGeneration.service";
import logger from "../config/logger.config";

export const setupRoomGenerationWorker = () => {

    const roomGenerationProcessor = new Worker<RoomGenerationJob>(
    ROOM_GENERATION_QUEUE, // Name of the queue
    async (job: Job) => {

        if(job.name !== ROOM_GENERATION_PAYLOAD){
            throw new Error("Invalid job name")
        }

        const payload = job.data
        logger.info(`Processing room generation for: ${JSON.stringify(payload)}`);

        await generateRoomsService(payload)
        logger.info(`Room generation completed for: ${JSON.stringify(payload)}`);


    },// Process function
    {
        connection: getRedisConnObject()
    }
)

roomGenerationProcessor.on('failed', () => {
    console.error('Room Generation processing failed');
})

roomGenerationProcessor.on('completed', (job) => {
    console.log('Room generation processing completed successfully for job:', job.name);
});
}