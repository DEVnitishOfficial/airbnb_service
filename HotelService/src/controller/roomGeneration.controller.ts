import { NextFunction, Request, Response } from "express";
import { StatusCodes } from "http-status-codes";
import { generateRoomsService } from "../services/roomGeneration.service";

export async function generateRoomsController(req: Request, res: Response, next: NextFunction){
    console.log('request recieved at controller', req.body);
    const result = await generateRoomsService(req.body);

    console.log("see the result from controller", result);
    
     res.status(StatusCodes.OK).json({
        success : true,
        message : "Rooms generated successfully",
        data : result
    })
}