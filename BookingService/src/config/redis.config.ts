import IORedis from 'ioredis';
import Redlock from 'redlock';
import { serverConfig } from '.';


export function connectToRedis() {
    try {
        let connection:IORedis;
        
        return ()=> {
            if(!connection){
                connection = new IORedis(serverConfig.REDIS_SERVER_URL);
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


export const redlock = new Redlock([getRedisConnObject()], {
    driftFactor: 0.01, // here the meaning of drift factor is the percentage of time that the lock can drift(extending lock duration to syncronize all the node suppose the lock TTL is 10 seconds and the driftFactor is 0.01 (1%), the extended lock duration will be 10.1 seconds. ) before it is considered expired
    retryCount: 10, // number of times to retry acquiring the lock
    retryDelay: 200, // here the meaning of retry delay is the time in ms to wait before retrying to acquire the lock
    retryJitter: 200,// retryJitter introduces randomness into the wait time, adding a random delay (up to the specified value in milliseconds) to the retryDelay. This randomness helps to prevent all clients from retrying at the same time

    // Example: If retryDelay is set to 200ms and retryJitter is also set to 200ms, a client will wait between 200ms and 400ms before retrying to acquire the lock after a failure
});