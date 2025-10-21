import { Job, Worker } from "bullmq";
import { HOTEL_INDEXING_QUEUE } from "../queues/hotelIndex.queue";
import { getRedisConnObject } from "../config/redis.config";
import { HOTEL_DELETION_PAYLOAD, HOTEL_INDEXING_PAYLOAD, HOTEL_UPDATE_PAYLOAD } from "../producers/hotelIndex.producer";
import { hotelIndexingHandlers } from "../utils/helpers/hotelIndexingHandlers";

const JOB_HANDLERS: Record<string, (data: any) => Promise<void>> = {
  // matching job names to their respective handler functions like
  // [JOB_NAME]           :           handlerFunction
  [HOTEL_INDEXING_PAYLOAD]: hotelIndexingHandlers.HOTEL_INDEXING_PAYLOAD,
  [HOTEL_UPDATE_PAYLOAD]: hotelIndexingHandlers.HOTEL_UPDATE_PAYLOAD,
  [HOTEL_DELETION_PAYLOAD]: hotelIndexingHandlers.HOTEL_DELETION_PAYLOAD,
};


const hotelIndexingProcessor = new Worker(
  HOTEL_INDEXING_QUEUE,
  async (job: Job) => {
    console.log(`Hotel indexing job received: ${job.name}`, job.data);

    const jobHandlerFn = JOB_HANDLERS[job.name];
    if (!jobHandlerFn) {
      console.warn(`No handler found for job: ${job.name}`);
      return;
    }

    await jobHandlerFn(job.data);
  },
  {
    connection: getRedisConnObject(),
  }
);

hotelIndexingProcessor.on("completed", (job) => {
  console.log(` Job completed successfully: ${job.name}`);
});

hotelIndexingProcessor.on("failed", (job, err) => {
  console.error(` Job ${job?.id} failed. Name: ${job?.name}\nError:`, err);
});

export default hotelIndexingProcessor;