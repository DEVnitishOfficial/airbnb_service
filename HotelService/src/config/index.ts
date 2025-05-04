
import dotenv from 'dotenv'

type ServerConfig = {
    PORT:number
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

export const serverConfig:ServerConfig = {
    PORT: Number(process.env.PORT) || 3002
}

export const dbConfig:DBconfig = {
    DB_USER:process.env.DB_USER || 'root',
    DB_PASSWORD:process.env.DB_PASSWORD || 'root',
    DB_NAME:process.env.DB_NAME || 'test_db',
    DB_HOST:process.env.DB_HOST || 'localhost'
}