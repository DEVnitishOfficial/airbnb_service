import { Queue } from "bullmq";
import { getRedisConnObject } from "../config/redis.config";

export const ROOM_GENERATION_QUEUE = 'room-generation-queue';

export const roomGenerationQueue = new Queue(ROOM_GENERATION_QUEUE, {
    connection: getRedisConnObject()
}); 
