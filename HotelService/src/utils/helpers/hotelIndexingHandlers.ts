import { ElasticsearchRepository } from "../../repositories/elasticsearch.repository";
import { HotelRepository } from "../../repositories/hotel.repository";
import { transformHotelToESDoc } from "../transformers/hotel.transformer";


const esRepo = new ElasticsearchRepository();
const hotelRepo = new HotelRepository();

export const hotelIndexingHandlers = {
  async HOTEL_INDEXING_PAYLOAD(jobData: any) {
    console.log("Processing hotel indexing job");

    const hotelId = jobData;
    const hotel = await hotelRepo.getHotelById(Number(hotelId));
    if (!hotel) {
      console.log(`Hotel with id ${hotelId} not found, skipping indexing.`);
      return;
    }

    const doc = transformHotelToESDoc(hotel);
    await esRepo.indexHotel(doc);
    console.log(`Hotel ${hotelId} indexed successfully in Elasticsearch`);
  },

  async HOTEL_UPDATE_PAYLOAD(jobData: any) {
    console.log("Processing hotel update job");

    const hotelId = jobData;
    const hotel = await hotelRepo.getHotelById(Number(hotelId));
    if (!hotel) {
      console.log(`Hotel with id ${hotelId} not found, skipping update.`);
      return;
    }

    const doc = transformHotelToESDoc(hotel);
    await esRepo.updateHotel(hotelId, doc);
    console.log(`Hotel ${hotelId} updated successfully in Elasticsearch`);
  },

  async HOTEL_DELETION_PAYLOAD(jobData: any) {
    console.log("Processing hotel deletion job");

    const hotelId = jobData;
    await esRepo.deleteHotel(hotelId);
    console.log(`Hotel ${hotelId} deleted successfully from Elasticsearch`);
  }
};
