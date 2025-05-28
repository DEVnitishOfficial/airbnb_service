import { Request, Response, NextFunction } from "express";
import { StatusCodes } from "http-status-codes";
import { confirmBookingService, createBookingService } from "../services/booking.service";

export async function  createBookingHandler(req: Request, res: Response, next: NextFunction) {
    // Call the service layer to create a booking
    const createBookingResponse = await createBookingService(req.body);
    // Send a success response with the created booking data
    res.status(StatusCodes.CREATED).json({
        success: true,
        message: "Booking created successfully",
        data: createBookingResponse,
    });
}

export async function confirmBookingHandler(req: Request, res: Response, next: NextFunction) {
    // Call the service layer to confirm a booking
    const booking = await confirmBookingService(req.params.idempotencyKey);
    // Send a success response with the confirmed booking data
    res.status(StatusCodes.OK).json({
        success: true,
        message: "Booking confirmed successfully",
        data: booking,
    });
}