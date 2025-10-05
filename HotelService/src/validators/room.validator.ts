import { z } from "zod";

export const getAvailableRoomSchema = z.object({
    roomCategoryId: z.number({message: "roomCategoryId required"}),
    checkInDate: z.string({message: "checkInDate required"}),
    checkOutDate: z.string({message: "checkOutDate required"}),
});

export const updateBookingIdToRoomsSchema = z.object({
    roomIds : z.array(z.number()).nonempty({message:"roomIds cannot be empty"}),
    bookingId: z.number({message:"bookingId required"})
})