
import dotenv from 'dotenv'

type ServerConfig = {
    PORT:number,
    REDIS_SERVER_URL:string,
    LOCK_TTL: number,
    DATABASE_URL?:string
}
export function loadEnv(){
    dotenv.config()
}

loadEnv()

export const serverConfig:ServerConfig = {
    PORT: Number(process.env.PORT) || 3002,
    REDIS_SERVER_URL: process.env.REDIS_SERVER_URL || 'redis://localhost:6379',
    LOCK_TTL: Number(process.env.LOCK_TTL) || 60000, // Default to 60 seconds
    DATABASE_URL: process.env.DATABASE_URL
}