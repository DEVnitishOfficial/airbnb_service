import { NextFunction, Request, Response } from "express";
import { allSoftDeletedHotelService, createHotelService, getAllHotelService, getHotelByIdService, hardDeleteHotelService, restoreSoftDeletedHotelByIdService, softDeleteHotelService, updateHotelService } from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";

export async function createHotelHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to create a hotel
        const hotelResponse = await createHotelService(req.body); 
        // Send a success response with the created hotel data
        res.status(StatusCodes.CREATED).json({  
            success: true,
            message: "Hotel created successfully",
            data: hotelResponse,
        });
}

export async function getHotelByIdHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to get a hotel by ID
        const hotelResponse = await getHotelByIdService(Number(req.params.id)); 
        // Send a success response with the hotel data
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "Hotel fetched successfully",
            data: hotelResponse,
        });
}

export async function getAllHotelsHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to get all hotels
        const allHotelResponse = await getAllHotelService();
        // Send a success response with the list of hotels
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "Hotels fetched successfully",
            data: allHotelResponse,
        });

}  

export async function updateHotelHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to update a hotel
        const updateHotelResponse = await updateHotelService(Number(req.params.id), req.body); 
        // Send a success response with the updated hotel data
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "Hotel updated successfully",
            data: updateHotelResponse,
        });
}

export async function softDeleteHotelHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to soft delete a hotel
        const softDeleteHotelResponse = await softDeleteHotelService(Number(req.params.id)); 
        // Send a success response indicating the hotel was soft deleted
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "Hotel soft deleted successfully",
            data: softDeleteHotelResponse,
        });

}

export async function hardDeleteHotelHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to hard delete a hotel
        const hardDeleteHotelResponse = await hardDeleteHotelService(Number(req.params.id)); 
        // Send a success response indicating the hotel was hard deleted
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "Hotel hard deleted successfully",
            data: hardDeleteHotelResponse,  
        });
} 

export async function allSoftDeleteHotelHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to hard delete a hotel
        const softDeletedHotelResponse = await allSoftDeletedHotelService(); 
        // Send a success response indicating the hotel was hard deleted
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "soft deleted hotels fetched successfully",
            data: softDeletedHotelResponse,  
        });
} 

export async function restoreSoftDeletedHotelByIdHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to hard delete a hotel
        const restoreSoftDeletedHotelResponse = await restoreSoftDeletedHotelByIdService(Number(req.params.id)); 
        // Send a success response indicating the hotel was hard deleted
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "restored soft deleted hotels successfully",
            data: restoreSoftDeletedHotelResponse,  
        });
}

