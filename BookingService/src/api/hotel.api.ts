import axios from "axios";
import { internalServerError, NotFoundError } from "../utils/errors/app.error";

const HOTEL_SERVICE_URL = process.env.HOTEL_SERVICE_URL || 'http://localhost:3001/api/v1/rooms/available';

export const getAvailableRooms = async ( roomCategoryId: number, checkInDate: Date, checkOutDate: Date) => {
    const response = await axios.post(HOTEL_SERVICE_URL, {
        
            roomCategoryId,
            checkInDate,
            checkOutDate
        
    });

    if (!response.data) {
        throw new internalServerError('Failed to fetch available rooms');
    }
    return response.data;
};

export const updateBookingIdToRoom = async (bookingId: number, roomIds: number[]) => {
    const UPDATE_BOOKING_URL = process.env.UPDATE_BOOKING_URL || 'http://localhost:3001/api/v1/rooms/update-booking-ids';

    const response = await axios.put(UPDATE_BOOKING_URL, {
        roomIds,
        bookingId
    });

    if (!response.data) {
        throw new internalServerError('Failed to update booking IDs');
    }
    return response.data;
};

// export const getAvailableRooms = async (hotelId: number, roomCategoryId: number, checkInDate: string, checkOutDate: string) => {
//     const response = await fetch(`http://localhost:8001/api/v1/rooms/available`, {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json'
//         },
//         body: JSON.stringify({
//             hotelId,
//             roomCategoryId,
//             checkInDate,
//             checkOutDate
//         })
//     });

//     if (!response.ok) {
//         throw new Error('Failed to fetch available rooms');
//     }

//     return response.json();
// };
