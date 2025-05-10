import { NextFunction, Request, Response } from "express";
import { createHotelService, getAllHotelService, getHotelByIdService } from "../services/hotel.service";

export async function createHotelHandler(req: Request, res: Response, next: NextFunction) {
    try {
        // Call the service layer to create a hotel
        const hotelResponse = await createHotelService(req.body); 
        // Send a success response with the created hotel data
        res.status(201).json({  
            success: true,
            message: "Hotel created successfully",
            data: hotelResponse,
        });
        
    } catch (error: any) {
        // Handle errors and send an error response
        res.status(500).json({ 
            success: false,
            message: "Something went wrong while creating the hotelService",
            Error: error.message,
            data: null,
        });
    }
}

export async function getHotelByIdHandler(req: Request, res: Response, next: NextFunction) {
    try {
        // Call the service layer to get a hotel by ID
        const hotelResponse = await getHotelByIdService(Number(req.params.id)); 
        // Send a success response with the hotel data
        res.status(200).json({  
            success: true,
            message: "Hotel fetched successfully",
            data: hotelResponse,
        });
        
    } catch (error: any) {
        // Handle errors and send an error response
        res.status(500).json({ 
            success: false,
            message: "Something went wrong while fetching the hotel",
            Error: error.message,
            data: null,
        });
    }
}

export async function getAllHotelsHandler(req: Request, res: Response, next: NextFunction) {
    try {
        // Call the service layer to get all hotels
        const allHotelResponse = await getAllHotelService();
        // Send a success response with the list of hotels
        res.status(200).json({  
            success: true,
            message: "Hotels fetched successfully",
            data: allHotelResponse,
        });
        
    } catch (error: any) {
        // Handle errors and send an error response
        res.status(500).json({ 
            success: false,
            message: "Something went wrong while fetching the hotels",
            Error: error.message,
            data: null,
        });
    }

}  

export async function updateHotelHandler(req: Request, res: Response, next: NextFunction) {
    res.status(501)
}

export async function deleteHotelHandler(req: Request, res: Response, next: NextFunction) {
    res.status(501)
}
export async function getHotelByLocationHandler(req: Request, res: Response, next: NextFunction) {
    res.status(501)
}

