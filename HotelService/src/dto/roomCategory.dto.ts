import { RoomType } from "../db/models/roomCategory";

export type CreateRoomCategoryDto = {
    hotelId: number;
    price: number;
    roomType : RoomType
    roomCount: number;
}