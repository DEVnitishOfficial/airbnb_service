// src/queues/hotelIndex.queue.ts
import { Queue } from "bullmq";
import { getRedisConnObject } from "../config/redis.config";

export const HOTEL_INDEXING_QUEUE = "hotel-indexing-queue";

export const HotelIndexQueue = new Queue(HOTEL_INDEXING_QUEUE, {
   connection: getRedisConnObject()
});
