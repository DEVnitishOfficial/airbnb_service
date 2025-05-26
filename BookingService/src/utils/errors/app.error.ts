export interface AppError extends Error{
    statusCode:number
}

export class internalServerError implements AppError{
    statusCode: number;
    message: string;
    name: string;

    constructor(message:string){
        this.statusCode = 500;
        this.name = "internalServerError";
        this.message = message;
    }
} 
export class NotFoundError implements AppError{
    statusCode: number;
    message: string;
    name: string;
    stack?: string | undefined;

    constructor(message:string){
        this.statusCode = 404;
        this.name = "NotFoundError";
        this.message = message;

        // Capture the stack trace
        if (Error.captureStackTrace) {
            Error.captureStackTrace(this, NotFoundError);
        } else {
            this.stack = new Error().stack;
        }
    }
} 
export class UnauthorizedError implements AppError{
    statusCode: number;
    message: string;
    name: string;

    constructor(message:string){
        this.statusCode = 401;
        this.name = "UnauthorizedError";
        this.message = message;
    }
} 
export class ConflictError implements AppError{
    statusCode: number;
    message: string;
    name: string;

    constructor(message:string){
        this.statusCode = 409;
        this.name = "ConflictError";
        this.message = message;
    }
} 
export class ValidatortError implements AppError{
    statusCode: number;
    message: string;
    name: string;

    constructor(message:string){
        this.statusCode = 422;
        this.name = "ValidatorError";
        this.message = message;
    }
} 

