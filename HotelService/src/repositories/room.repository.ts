import Room from "../db/models/room";
import BaseRepository from "./base.repository";


class RoomRepository extends BaseRepository<Room> {
    constructor() {
        super(Room); // Assuming "Room" is the name of the model
    }
}

export default RoomRepository;