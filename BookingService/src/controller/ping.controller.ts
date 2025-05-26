import { NextFunction, Request, Response } from "express"
import fs from 'fs/promises'
import {NotFoundError } from "../utils/errors/app.error"
import logger from "../config/logger.config"
export const pingHandler =async (req:Request,res:Response,next:NextFunction) => {

    try{
        // await fs.readFile("sample")
        res.status(200).json({message:"pong"})
        logger.info('message received successfully')
    }catch(error){
        logger.error('something went wrong read the error carefully',error);
        throw new NotFoundError("File not found....")
    }
    
}
