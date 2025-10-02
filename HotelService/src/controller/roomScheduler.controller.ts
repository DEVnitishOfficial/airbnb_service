
import { Request, Response } from "express";
import { startRoomSchedulerCronJob } from "../scheduler/roomScheduler";
import { StatusCodes } from "http-status-codes";

export async function startRoomSchedulerController(req:Request, res:Response){

        startRoomSchedulerCronJob()

        res.status(StatusCodes.OK).json({
            message: "Room availability extension scheduler started successfully",
            success: true,
            data: {
                status: "started"
            }
        });
}