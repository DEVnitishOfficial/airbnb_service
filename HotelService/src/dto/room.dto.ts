
export type GetAvailableRoomsDTO = {
    roomCategoryId: number;
    checkInDate: Date;
    checkOutDate: Date;
}

export type UpdateBookingIdToRoomsDTO = {
    roomIds: number[];
    bookingId: number;
}