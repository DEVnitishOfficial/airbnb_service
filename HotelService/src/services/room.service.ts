import { GetAvailableRoomsDTO, UpdateBookingIdToRoomsDTO } from "../dto/room.dto";
import RoomRepository from "../repositories/room.repository";


const roomRepository = new RoomRepository();

export async function getAvailableRoomsService(getAvailableRoomsDTO: GetAvailableRoomsDTO){
    const { roomCategoryId, checkInDate, checkOutDate } = getAvailableRoomsDTO;
    const rooms = await roomRepository.findByRoomCategoryIdAndDateRange(roomCategoryId, checkInDate, checkOutDate);
    return rooms;
}

export async function updateBookingIdToRoomsService(updateBookingIdToRoomsDTO: UpdateBookingIdToRoomsDTO){
    const { roomIds, bookingId } = updateBookingIdToRoomsDTO;
    return await roomRepository.updateBookingIdToRooms(roomIds, bookingId);
}
