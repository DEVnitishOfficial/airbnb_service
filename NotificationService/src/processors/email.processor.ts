import { Job, Worker } from "bullmq";
import { NotificationDTO } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { EMAIL_PRODUCER } from "../producers/email.producer";

export const setupEmailWorker = () => {
    const emailProcessor = new Worker<NotificationDTO>(
    MAILER_QUEUE, // Name of the queue
    async (job: Job) => {
        if (job.name !== EMAIL_PRODUCER) {
            throw new Error(`Invalid job name: ${job.name}`);
        }

        //call service layer from here to process the email
        const payload = job.data;
        console.log(`Processing email :  ${payload.to} with subject: ${payload.subject}`);



    },// Process function
    {
        connection: getRedisConnObject()
    }
)

emailProcessor.on('failed', () => {
    console.error('Email processing failed');
})

emailProcessor.on('completed', (job) => {
    console.log('Email processing completed successfully for job:', job.name);
});
}