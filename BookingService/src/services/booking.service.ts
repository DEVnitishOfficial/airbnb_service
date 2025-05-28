import { CreateBookingDTO } from "../dto/booking.dto";
import { confirmBooking, createBooking, createIdempotencyKey, finalizeIdempotencyKey, getIdemPotencyKey } from "../repositories/booking.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/helpers/generateIdempotencyKey";

export async function createBookingService( createBookingDTO: CreateBookingDTO){
    const booking = await createBooking({
        userId : createBookingDTO.userId,
        hotelId: createBookingDTO.hotelId,
        bookingAmount: createBookingDTO.bookingAmount,
        totalGuests: createBookingDTO.totalGuests,
    });

    const idempotencyKey = generateIdempotencyKey();
    await createIdempotencyKey(idempotencyKey, booking.id);

    return {
        bookingId: booking.id,
        idempotencyKey: idempotencyKey
    }
}
export async function confirmBookingService(idempotencyKey: string) {
    // This function will be used to finalize the booking after payment is confirmed
    const idempotencyKeyData = await getIdemPotencyKey(idempotencyKey);
    if (!idempotencyKeyData) {
        throw new NotFoundError("Idempotency key not found");
    }
    if (idempotencyKeyData.finalized) {
        throw new BadRequestError("Booking already finalized");
    }
    const booking = await confirmBooking(idempotencyKeyData.bookingId);
    await finalizeIdempotencyKey(idempotencyKey);
    return booking;
}