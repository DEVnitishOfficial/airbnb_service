import Redis from "ioredis";
import { serverConfig } from ".";

// export function connectToRedis() {
//     try {
//         const redisConfig = {
//             host: serverConfig.REDIS_HOST,
//             port: serverConfig.REDIS_PORT,
//         }
//         const connection = new Redis(redisConfig);
//         return connection;

//     } catch (error) {
//         console.error("Error connecting to Redis:", error);
//         throw error;
//     }
// }

// Here i have used singelton pattern to ensure only one connection is created throughout the application lifecycle
// This is useful for performance and resource management
export function connectToRedis() {
    try {
        let connection:Redis;
        const redisConfig = {
            host: serverConfig.REDIS_HOST,
            port: serverConfig.REDIS_PORT,
            maxRetriesPerRequest: null, // Disable automatic reconnection
        }
        
        return ()=> {
            if(!connection){
                connection = new Redis(redisConfig);
                return connection;
            }
            return connection;
        }
       

    } catch (error) {
        console.error("Error connecting to Redis:", error);
        throw error;
    }
}

export const getRedisConnObject = connectToRedis();