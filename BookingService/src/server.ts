import express from 'express'
import { serverConfig } from './config';
import v1Router from './routers/v1/index.router';
import v2Router from './routers/v2/index.router';
import { genericErrorHandler } from './middlewares/error.middleware';
import logger from './config/logger.config';
import { attachCorrelationIdMiddleware } from './middlewares/correlation.middleware';
import { addEmailToQueue } from './producer/email.producer';



const app = express();
app.use(express.json());


app.use(attachCorrelationIdMiddleware)

app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);

app.use(genericErrorHandler);

app.listen(serverConfig.PORT, () => {
    console.log(`server is listening at  http://localhost:${serverConfig.PORT}`)
    logger.info("press Ctrl + C to stop the server", { "name": "dev-server" })

    for(let i=0; i<10; i++){
        const sampleNotification = {
            to: `sample${i} from bookingservice`,
            subject: `Sample Subject ${i} bookin service`,
            templateId: `sample-template ${i} booking service`,
            params: {
                name: `Nitish official ${i}`,
                orderId: `${12345 + i}`,
                message: `This is a sample notification message from booking service. ${i}`
            }
        };
        addEmailToQueue(sampleNotification);
    }
});