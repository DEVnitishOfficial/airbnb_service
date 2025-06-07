import nodemailer from 'nodemailer'
import { serverConfig } from '.';

const transpoter = nodemailer.createTransport({
        service:'gmail',
        auth : {
            user : serverConfig.MAIL_USER,
            pass : serverConfig.MAIL_PASS
        }
});


export default transpoter;


