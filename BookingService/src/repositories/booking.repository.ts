import { Booking, Prisma } from "@prisma/client"
import  PrismaClient  from "../prisma/client";

export async function createBooking(bookingInput : Prisma.BookingCreateInput) {
    const booking = await PrismaClient.booking.create({
        data: bookingInput
    })
    return booking;
}

export async function createIdempotencyKey(key: string, bookingId: number) {
    const idempotencyKey = await PrismaClient.idempotencyKey.create({
        data: {
            key, // this will be a number
            booking: {
                connect: { // here connect method is used to connect the booking with the idempotency key
                    id: bookingId
                }
            }
        }
    })
    return idempotencyKey;
}

export async function getIdemPotencyKey(key: string) {
    const idempotencyKey = await PrismaClient.idempotencyKey.findUnique({
        where: {
            key
        }
    })
    return idempotencyKey;
}   

export async function getBookingById(bookingId: number) {
    const booking = await PrismaClient.booking.findUnique({
        where: {
            id : bookingId
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

export async function confirmBooking(bookingId: number) {
    const booking = await PrismaClient.booking.update({
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

export async function finalizeIdempotencyKey(key: string) {
    const idempotencyKey = await PrismaClient.idempotencyKey.update({
        where: {
            key
        },
        data: {
            finalized: true
        }
    })
    return idempotencyKey;
}