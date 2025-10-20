import { HotelRecord } from "../../dto/hotel.dto";

export function transformHotelToESDoc(hotelRecord: HotelRecord) {
  console.log('hotelData received for transormation>>>>', hotelRecord);
  return {
    id: String(hotelRecord.id),
    name: hotelRecord.name,
    address: hotelRecord.address,
    location: hotelRecord.location
  };
}
