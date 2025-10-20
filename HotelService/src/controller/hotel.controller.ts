import { NextFunction, Request, Response } from "express";
import {createHotelService, deleteHotelService, getAllHotelService, getHotelByIdService,updateHotelService } from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";
import { addHotelIndexInESJobToQueue, deleteHotelIndexInESJobToQueue } from "../producers/hotelIndex.producer";

export async function createHotelHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to create a hotel
        const hotelResponse = await createHotelService(req.body); 

        // Enqueue a job to index the newly created hotel in Elasticsearch
        await addHotelIndexInESJobToQueue(hotelResponse.id);
        
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

        // Enqueue a job to update the hotel in Elasticsearch
        await addHotelIndexInESJobToQueue(Number(req.params.id));
        
        // Send a success response with the updated hotel data
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "Hotel updated successfully",
            data: updateHotelResponse,
        });
}

export async function deleteHotelHandler(req: Request, res: Response, next: NextFunction) {
        // Call the service layer to delete a hotel
        await deleteHotelService(Number(req.params.id));

        // Enqueue a job to delete the hotel from Elasticsearch
        await deleteHotelIndexInESJobToQueue(Number(req.params.id));

        // Send a success response indicating the hotel was deleted
        res.status(StatusCodes.OK).json({  
            success: true,
            message: "Hotel deleted successfully",
        });
}
