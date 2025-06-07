import { serverConfig } from "../config";
import logger from "../config/logger.config";
import transpoter from "../config/mailer.config";
import { internalServerError } from "../utils/errors/app.error";

    

export async function sendEmail(to:string,subject:string,body:string){
    try {
        await transpoter.sendMail({
            from : serverConfig.MAIL_USER,
            to,
            subject,
            html : body
        })
        logger.info(`Email sent to :${to} with subject : ${subject}`)
    } catch (error) {
        throw new internalServerError(`Failed to send mail : ${error}`)
    }
}