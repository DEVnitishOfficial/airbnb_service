
import winston from "winston"
import { getCorrelationId } from "../utils/helpers/request.helpers";
import DailyRotateFile from 'winston-daily-rotate-file'
import { MongoDB } from "winston-mongodb";

const logger = winston.createLogger({
    format: winston.format.combine(
        winston.format.timestamp({ format: "MM:DD:YYYY HH:MM:SS" }),
        winston.format.json(), // format the log message as json
        // defining custom print
        winston.format.printf(({ timestamp, level, message, ...data }) => {
            const output = { level, message, timestamp, correlationId: getCorrelationId(), data };
            return JSON.stringify(output)
        })
    ),
    transports:[
        new winston.transports.Console(),
        // new winston.transports.File({filename:'logs/app.log'})
        new DailyRotateFile({
        filename: 'logs/application-%DATE%-app.log',
        datePattern: 'YYYY-MM-DD',
        maxSize: '20m',
        maxFiles: '14d'
        }),

       

        // TODO: add logic to integrate and save logs in mongodb
        // new MongoDB({
        //     db: String(process.env.MONGO_URI) || 'mongodb://localhost:27017/logs_db',
        //     options: {
        //       useUnifiedTopology: true,
        //     },
        //     collection: 'log_entries',
        //     tryReconnect: true,
        //   })
        
    ] 
})

export default logger;