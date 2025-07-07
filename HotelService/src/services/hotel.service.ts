import { createHotelDto, updateHotelDto } from "../dto/hotel.dto";
import { HotelRepository} from "../repositories/hotel.repository"
import { BadRequesError } from "../utils/errors/app.error";


const hotelRepository = new HotelRepository();

export async function createHotelService(hotelData: createHotelDto) {
    const hotel = await hotelRepository.create(hotelData);
    return hotel;
}

export async function getHotelByIdService(id: number) {
    const hotel = await hotelRepository.findById(id);
    return hotel;
}

export async function getAllHotelService() {
    const hotels = await hotelRepository.findAll();
    return hotels;
}

export async function updateHotelService(id: number, hotelData: updateHotelDto) {
    const updatedHotel = await hotelRepository.update(id, hotelData);
    return updatedHotel;
}

export async function deleteHotelService(id: number) {
    const hotel = await hotelRepository.findById(id);
    if (!hotel) {
        throw new BadRequesError(`Hotel with id ${id} not found`);
    }
    await hotelRepository.delete({ id });
    return;
}
