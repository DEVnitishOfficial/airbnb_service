
import dotenv from 'dotenv'

type ServerConfig = {
    PORT:number,
    REDIS_HOST: string,
    REDIS_PORT: number,
    MAIL_USER:string,
    MAIL_PASS:string,
}
function loadEnv(){
    dotenv.config()
}

loadEnv()

export const serverConfig:ServerConfig = {
    PORT: Number(process.env.PORT) || 3002,
    REDIS_HOST: process.env.REDIS_HOST || 'localhost',
    REDIS_PORT: Number(process.env.REDIS_PORT) || 6379,
    MAIL_USER: process.env.MAIL_USER || '',
    MAIL_PASS: process.env.MAIL_PASS || ''
}