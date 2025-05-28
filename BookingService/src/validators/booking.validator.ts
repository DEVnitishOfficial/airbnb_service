import { z } from "zod";

export const bookingSchema = z.object({
    // Define the schema for hotel creation
    userId: z.number({message : "userId required"}),
    hotelId: z.number({message : "hotelId required"}),
    totalGuests: z.number({message: "totalGuests required"}).min(1, {message: "totalGuests must be at least 1"}),
    bookingAmount: z.number({message: "bookingAmount required"}).min(1, {message: "bookingAmount must be at least 1"}),
});