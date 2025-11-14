import { Job, Worker } from "bullmq";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import { ROOM_GENERATION_QUEUE } from "../queues/roomGeneration.queue";
import { getRedisConnObject } from "../config/redis.config";
import { ROOM_GENERATION_PAYLOAD } from "../producers/roomGeneration.producer";
import { generateRoomsService } from "../services/roomGeneration.service";
import logger from "../config/logger.config";
import { asyncLocalStorage } from "../utils/helpers/request.helpers";

export const setupRoomGenerationWorker = () => {

    const roomGenerationProcessor = new Worker<RoomGenerationJob & { correlationId: string }>(
        ROOM_GENERATION_QUEUE, // Name of the queue
        async (job: Job) => {

            if (job.name !== ROOM_GENERATION_PAYLOAD) {
                throw new Error("Invalid job name")
            }

            const payload = job.data
            const correlationId = job.data.correlationId
            logger.info(`Processing room generation for: ${JSON.stringify(payload)} and correlationId: ${correlationId}`);
            // --- CONTEXT RE-ESTABLISHMENT ---
            // 1. Run the entire job processor logic inside the ALS.run() wrapper
            // This makes the correlationId available to all downstream functions (logger, services, etc.)
            return asyncLocalStorage.run({ correlationId }, async () => {

                logger.info(`Processing room generation for: ${JSON.stringify(payload)}`);

                // Example: logger.info will now automatically include the correlationId
                await generateRoomsService(payload);

                logger.info(`Room generation completed for: ${JSON.stringify(payload)}`);
            });
            // --- END RE-ESTABLISHMENT ---


        },
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