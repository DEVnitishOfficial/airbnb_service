
import dotenv from 'dotenv'

type ServerConfig = {
    PORT:number,
    REDIS_HOST: string,
    REDIS_PORT: number,
    ROOM_CRON: string,
}

type DBconfig = {
DB_USER:string,
DB_PASSWORD:string,
DB_NAME:string,
DB_HOST:string

}

function loadEnv(){
    dotenv.config()
}

loadEnv()

console.log('see my env>>>>',process.env.PORT)

export const serverConfig:ServerConfig = {
    PORT: Number(process.env.PORT) || 3002,
    REDIS_HOST: process.env.REDIS_HOST || 'localhost',
    REDIS_PORT: Number(process.env.REDIS_PORT) || 6379,
    //  ROOM_CRON: process.env.ROOM_CRON || '0 2 * * *', // Run the job every day at 2:00 AM.
    ROOM_CRON: process.env.ROOM_CRON || '* * * * *', // Run the job every minute.
}

export const dbConfig:DBconfig = {
    DB_USER:process.env.DB_USER || 'root',
    DB_PASSWORD:process.env.DB_PASSWORD || 'root',
    DB_NAME:process.env.DB_NAME || 'test_db',
    DB_HOST:process.env.DB_HOST || 'localhost',
}