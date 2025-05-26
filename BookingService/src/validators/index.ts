import { NextFunction, Request,Response } from "express"
import { AnyZodObject } from "zod"
import logger from "../config/logger.config"


export const validateRequestBody = (schema:AnyZodObject) => {
    return async (req:Request,res:Response,next:NextFunction)=> {
        try{
            logger.info("validation the request body log-first")
            await schema.parseAsync(req.body)
            logger.info('request body is validated succcessfully log-second');
        }catch(error){
            logger.error("request body is invalid")
            res.status(400).json({
            success:false,
            message: "invalid schema",
            error:error
           })
        }
        next();
    }
}

export const validateQueryParams = (schema:AnyZodObject) => {
    return async (req:Request,res:Response,next:NextFunction)=> {
        try{
            await schema.parseAsync(req.query)
            console.log('Query params is validated succcessfully');
        }catch(error){
            res.status(400).json({
            success:false,
            message: "invalid schema",
            error:error
           })
        }
    }
}