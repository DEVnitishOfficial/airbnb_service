import { Prisma, IdempotencyKey } from "@prisma/client"
import { validate as isValidUUID } from "uuid";
import PrismaClient from "../prisma/client";
import { BadRequestError } from "../utils/errors/app.error";

export async function createBooking(bookingInput: Prisma.BookingCreateInput) {
    const booking = await PrismaClient.booking.create({
        data: bookingInput
    })
    return booking;
}

export async function createIdempotencyKey(key: string, bookingId: number) {
    const idempotencyKey = await PrismaClient.idempotencyKey.create({
        // here create method create a idempotency key and connect to the existing booking
        data: {
            idemKey: key, // this is uuid 
            booking: {
                connect: { // here connect method is used to connect the booking with the idempotency key
                    id: bookingId
                }
            }
        }
    })
    return idempotencyKey;
}

export async function getIdemPotencyKeyWithLock(tx: Prisma.TransactionClient, key: string) {

    if (!isValidUUID(key)) {
        throw new BadRequestError("Invalid idempotency key format");
    }

    // This function is used to get the idempotency key with a lock
    const idempotencyKey: Array<IdempotencyKey> = await tx.$queryRaw(
        Prisma.raw(`SELECT * FROM IdempotencyKey WHERE idemKey = '${key}' FOR UPDATE;`))

    if (!idempotencyKey || idempotencyKey.length === 0) {
        throw new BadRequestError("Idempotency key not found");
    }
    // here we are returning the first element of the array because the query will return an array of idempotency keys because we are using select * and put some conditions so it check the matching key and return the all matching keys, but in our case we are using unique key so it will return only one key, that's why we are returning the first element of the array
    return idempotencyKey[0];
}

export async function getBookingById(bookingId: number) {
    const booking = await PrismaClient.booking.findUnique({
        where: {
            id: bookingId
        }
    })
    return booking;
}

// export async function changeBookingStatus(bookingId: number, status: Prisma.EnumBookingStatusFieldUpdateOperationsInput) {
//     const booking = await PrismaClient.booking.update({
//         where: {
//             id: bookingId
//         },
//         data: {
//             status: status
//         }
//     })
//     return booking;
// }

export async function confirmBooking(tx: Prisma.TransactionClient, bookingId: number) {
    const booking = await tx.booking.update({
        where: {
            id: bookingId
        },
        data: {
            status: "CONFIRMED"
        }
    })
    return booking;
}

export async function cancelBooking(bookingId: number) {
    const booking = await PrismaClient.booking.update({
        where: {
            id: bookingId
        },
        data: {
            status: "CANCELLED"
        }
    })
    return booking;
}

export async function finalizeIdempotencyKey(tx: Prisma.TransactionClient, key: string) {
    const idempotencyKey = await tx.idempotencyKey.update({
        where: {
            idemKey: key
        },
        data: {
            finalized: true
        }
    })
    return idempotencyKey;
}