import { Job, Worker } from "bullmq";
import { HOTEL_INDEXING_QUEUE } from "../queues/hotelIndex.queue";
import { ElasticsearchRepository } from "../repositories/elasticsearch.repository";
import { getRedisConnObject } from "../config/redis.config";
import { transformHotelToESDoc } from "../utils/transformers/hotel.transformer";
import { HotelRepository } from "../repositories/hotel.repository";
import { HOTEL_DELETION_PAYLOAD, HOTEL_INDEXING_PAYLOAD, HOTEL_UPDATE_PAYLOAD } from "../producers/hotelIndex.producer";

const hotelIndexingProcessor = new Worker(HOTEL_INDEXING_QUEUE, async (job: Job) => {

  console.log("Hotel indexing job received:", job.name, job.data);

  const esRepo = new ElasticsearchRepository();
  const hotelRepo = new HotelRepository();

  if (job.name === HOTEL_INDEXING_PAYLOAD) {
    console.log("Processing hotel indexing job");
    const hotelId = job.data;
    const hotel = await hotelRepo.getHotelById(Number(hotelId));

    console.log("log the coming hotel response>>>>", hotel);

    if (!hotel) {
      console.log(`Hotel with id ${hotelId} not found, skipping indexing.`);
      return;
    }

    const doc = transformHotelToESDoc(hotel);
    await esRepo.indexHotel(doc);
    console.log(`Hotel ${hotelId} indexed successfully in Elasticsearch`);

  } else if (job.name === HOTEL_UPDATE_PAYLOAD) {
    console.log("Processing hotel update job");
    const hotelId = job.data;
    const hotel = await hotelRepo.getHotelById(Number(hotelId));
    if (!hotel) {
      console.log(`Hotel with id ${hotelId} not found, skipping update.`);
      return;
    }
    const doc = transformHotelToESDoc(hotel);
    await esRepo.updateHotel(hotelId, doc);
    console.log(`Hotel ${hotelId} updated successfully in Elasticsearch`);


  } else if (job.name === HOTEL_DELETION_PAYLOAD) {
    console.log("Processing hotel deletion job");
    const hotelId = job.data;
    await esRepo.deleteHotel(hotelId);
  }
}, {
  connection: getRedisConnObject()
});

hotelIndexingProcessor.on("completed", (job) => {
  console.log("Hotel index job completed successfully for job:", job.name);
});

hotelIndexingProcessor.on("failed", (job, err) => {
  console.error(`Job ${job?.id} failed. Name: ${job?.name}
Error:`, err);
});


export default hotelIndexingProcessor;
