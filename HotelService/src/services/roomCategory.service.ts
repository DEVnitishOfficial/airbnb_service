import { CreateRoomCategoryDto } from "../dto/roomCategory.dto";
import { HotelRepository } from "../repositories/hotel.repository";
import RoomCategoryRepository from "../repositories/roomCategory.repository";
import { NotFoundError } from "../utils/errors/app.error";


const roomCategoryRepository = new RoomCategoryRepository();
const hotelRepository = new HotelRepository();

export async function createRoomCategoryService(roomCategoryData: CreateRoomCategoryDto) {
    const roomCategory = await roomCategoryRepository.create(roomCategoryData);
    return roomCategory;
}

export async function getRoomCategoryByIdService(id: number) {
    const roomCategory = await roomCategoryRepository.findById(id);
    return roomCategory;
}

export async function getAllRoomCategoriesByHotelIdService(hotelId: number) {
    // check if hotel exists
    const hotel = await hotelRepository.findById(hotelId);
    if (!hotel) {
        throw new NotFoundError(`Hotel with id ${hotelId} does not exist`);
    }
    // find all room categories by hotel id
    const roomCategories = await roomCategoryRepository.findAllByHotelId(hotelId);
    return roomCategories;
}

export async function deleteRoomCategoryByIdService(id: number) {

    const roomCategory = await roomCategoryRepository.findById(id);
    if (!roomCategory) {
        throw new NotFoundError(`Room category with id ${id} does not exist`);
    }
    await roomCategoryRepository.delete({ id });
    return roomCategory;
}

