import { Request, Response } from "express";
import { deleteBookingIdFromRoomsService, getAvailableRoomsService, updateBookingIdToRoomsService } from "../services/room.service";
import { StatusCodes } from "http-status-codes";


export async function getAvailableRoomsController(req: Request, res: Response) {

    const availableRooms = await getAvailableRoomsService(req.body);
    res.status(StatusCodes.OK).json({
        success: true,
        message: "Available rooms fetched successfully",
        data: availableRooms
    });
}

export async function updateBookingIdToRoomsController(req: Request, res: Response) {
    const result = await updateBookingIdToRoomsService(req.body);
    res.status(StatusCodes.OK).json({
        success: true,
        message: "Booking IDs updated to rooms table successfully",
        data: result
    });
}

export async function releaseRoomsController(req: Request, res: Response) {
    const { bookingId } = req.body;
    const result = await deleteBookingIdFromRoomsService(bookingId);
    res.status(StatusCodes.OK).json({
        success: true,
        message: "Booking ID removed from rooms table successfully",
        data: result
    });
}