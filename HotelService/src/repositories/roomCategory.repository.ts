import RoomCategory from "../db/models/roomCategory";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repository";


 class RoomCategoryRepository extends BaseRepository<RoomCategory>{
    constructor() {
        super(RoomCategory);
    }

    async findAllByHotelId(hotelId: number): Promise<RoomCategory[]> {
        const roomCategories = await this.model.findAll({
            where: { 
                hotelId: hotelId,
                deletedAt: null, // Ensure we only get non-deleted room categories
            },
        });

        if (!roomCategories || roomCategories.length === 0) {
            throw new NotFoundError(`No room categories found for hotel with id ${hotelId}`);
        }

        return roomCategories;
    }
}

export default RoomCategoryRepository;
