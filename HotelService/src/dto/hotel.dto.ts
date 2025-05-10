export type createHotelDto = {
    name: string;
    address: string;
    location: string;
    rating?: number;
    ratingCount?: number;
}

export type updateHotelDto = {
    name?: string;
    address?: string;
    location?: string;
    rating?: number;
    ratingCount?: number;
}