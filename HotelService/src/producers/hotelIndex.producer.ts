import { HotelIndexQueue } from "../queues/hotelIndex.queue";

console.log("Hotel Index Queue initialized");

export const HOTEL_INDEXING_PAYLOAD = "payload:hotel-indexing";
export const HOTEL_DELETION_PAYLOAD = "payload:hotel-deletion";

export async function addHotelIndexInESJobToQueue(hotelId: number) {
  console.log(`Adding hotel indexing job to queue for hotel ID: ${hotelId}`);
  await HotelIndexQueue.add(HOTEL_INDEXING_PAYLOAD, hotelId);
}

export async function deleteHotelIndexInESJobToQueue(hotelId: number) {
  console.log(`Adding hotel deletion job to queue for hotel ID: ${hotelId}`);
  await HotelIndexQueue.add(HOTEL_DELETION_PAYLOAD, hotelId);
}
