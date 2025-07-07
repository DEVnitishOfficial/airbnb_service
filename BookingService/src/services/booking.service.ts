import { CreateBookingDTO } from "../dto/booking.dto";
import { confirmBooking, createBooking, createIdempotencyKey, finalizeIdempotencyKey, getIdemPotencyKeyWithLock } from "../repositories/booking.repository";
import { BadRequestError, internalServerError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/helpers/generateIdempotencyKey";
import PrismaClient from "../prisma/client";
import { serverConfig } from "../config";
import { redlock } from "../config/redis.config";

export async function createBookingService(createBookingDTO: CreateBookingDTO) {
    const ttl = serverConfig.LOCK_TTL;
    const bookingResource = `booking:${createBookingDTO.hotelId}`;

    const booking = await PrismaClient.booking.findFirst({
        where: {
            userId: createBookingDTO.userId,
            hotelId: createBookingDTO.hotelId,
        }
    })

    if (booking) {
        throw new BadRequestError(`You have already created booking with the same userId ${createBookingDTO.userId} and hotelId : ${createBookingDTO.hotelId}`)
    } else {
        try {
            await redlock.acquire([bookingResource], ttl); // here redlock takes two parameters, the first is an array of resources to lock, and the second is the TTL for the lock in milliseconds.
            const booking = await createBooking({
                userId: createBookingDTO.userId,
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

        } catch (error) {
            throw new internalServerError("Failed to acquire lock for booking resource");
        }
    }


}

export async function confirmBookingService(idempotencyKey: string) {

    return await PrismaClient.$transaction(async (tx) => {
        // Check if the idempotency key exists and is not finalized
        const idempotencyKeyData = await getIdemPotencyKeyWithLock(tx, idempotencyKey);
        if (!idempotencyKeyData) {
            throw new NotFoundError("Idempotency key not found");
        }
        if (idempotencyKeyData.finalized) {
            throw new BadRequestError("Booking already finalized");
        }
        const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
        await finalizeIdempotencyKey(tx, idempotencyKey);
        return booking;
    });
}