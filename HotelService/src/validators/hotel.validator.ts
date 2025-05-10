import { z } from "zod";

export const hotelSchema = z.object({
    // Define the schema for hotel creation
    name: z.string().min(1),
    address: z.string().min(1),
    location: z.string().min(1),
    rating: z.number().optional(),
    ratingCount: z.number().optional(),
});
