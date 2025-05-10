import { createHotelDto, updateHotelDto } from "../dto/hotel.dto";
import { createHotel, deleteHotelById, getAllHotel, getHotelById, updateHotelById } from "../repositories/hotel.repository"
import { BadRequesError } from "../utils/errors/app.error";

/**  const blockListedAddresses = [
    "123 Main St",
    "456 Elm St",
    "789 Oak St"
];

export function isAddressBlockListed(address: string): boolean {
    return blockListedAddresses.includes(address);
}

export async function createHotelServiceDemo(hotelData: createHotelDto) {
    if (isAddressBlockListed(hotelData.address)) {
        throw new BadRequesError("Address is blocklisted");
    }
    const hotel = await createHotel(hotelData);
    return hotel;
}
*/

// The above commented is a business logic that checks if the address is blocklisted then insted of creating the hotel it will throw an error.


// The createHotelService function is a service layer function that creates a new hotel.
// It calls the createHotel function from the repository layer and returns the created hotel object.
export async function createHotelService(hotelData: createHotelDto) {
    const hotel = await createHotel(hotelData);
    return hotel;
}
// The getHotelByIdService function is a service layer function that retrieves a hotel by its ID.
// It calls the getHotelById function from the repository layer and returns the hotel object.
export async function getHotelByIdService(id: number) {
    const hotel = await getHotelById(id);
    return hotel;
}

export async function getAllHotelService() {
    const hotels = await getAllHotel();
    return hotels;
}

export async function updateHotelService(id: number, hotelData: updateHotelDto) {
    const updatedHotel = await updateHotelById(id, hotelData);
    return updatedHotel;
}

export async function deleteHotelService(id: number) {
    const deletedHotel = await deleteHotelById(id);
    return deletedHotel;
}