import { NextFunction, Request, Response } from "express";
import { StatusCodes } from "http-status-codes";
import { addRoomGenerationJobToQueue } from "../producers/roomGeneration.producer";
import { getCorrelationId } from "../utils/helpers/request.helpers";

export async function generateRoomsController(req: Request, res: Response){
    console.log('request recieved at controller', req.body);

     const currentCorrelationId = getCorrelationId();

    // const result = await generateRoomsService(req.body);
    await addRoomGenerationJobToQueue(req.body, currentCorrelationId)

     res.status(StatusCodes.OK).json({
        success : true,
        message : "Added room generation job to Queue successfully",
        data : {}
    })
}