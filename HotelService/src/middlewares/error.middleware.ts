import { NextFunction, Request, Response } from "express";
import { AppError } from "../utils/errors/app.error";

export const genericErrorHandler = (err: AppError, req: Request, res: Response, next: NextFunction) => {
    console.log('eror from generic error', err);
    res.status(err.statusCode || 500).json({ success: false, message: err.message, stack:err.stack })
}