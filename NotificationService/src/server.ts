import express from 'express'
import { serverConfig } from './config';
import v1Router from './routers/v1/index.router';
import v2Router from './routers/v2/index.router';
import { genericErrorHandler } from './middlewares/error.middleware';
import logger from './config/logger.config';
import { attachCorrelationIdMiddleware } from './middlewares/correlation.middleware';
import { setupEmailWorker } from './processors/email.processor';
import { NotificationDTO } from './dto/notification.dto';
import { addEmailToQueue } from './producers/email.producer';


const app = express();
app.use(express.json());


app.use(attachCorrelationIdMiddleware)

app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);

app.use(genericErrorHandler);

app.listen(serverConfig.PORT, () => {
    console.log(`server is listening at  http://localhost:${serverConfig.PORT}`)
    logger.info("press Ctrl + C to stop the server", { "name": "dev-server" })
    setupEmailWorker();
    logger.info("Email worker setup completed");

    // const sampleNotification : NotificationDTO = {
    //     to: "sample",
    //     subject: "Sample Subject",
    //     templateId: "sample-template",
    //     params: {
    //         name: "Nitish official",
    //         orderId: "12345",
    //         message: "This is a sample notification message."

    //     }
    // };
    // addEmailToQueue(sampleNotification)
    logger.info("Sample email added to queue");
})