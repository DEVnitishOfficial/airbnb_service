import express from "express"
import { pingHandler } from "../../controller/ping.controller"
import { validateRequestBody } from "../../validators"
import { pingSchema } from "../../validators/ping.validator"

const pingRouter = express.Router()

function checkHandler(req:express.Request,res:express.Response,next:express.NextFunction):void {
    if(typeof req.body.name !== 'string'){
       res.status(400).json({
            success:false,
            message:"something went wrong"
        })
    }
    next()
    }


pingRouter.get('/',validateRequestBody(pingSchema), pingHandler);



pingRouter.get('/health',(req,res)=> {
    res.send('Your health is good')
})

export default pingRouter;