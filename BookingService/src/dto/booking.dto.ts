export type CreateBookingDTO = {
    userId: number
    hotelId: number;
    bookingAmount: number;
    totalGuests: number;
    checkInDate: Date;
    checkOutDate: Date;
    roomCategoryId: number;
}
