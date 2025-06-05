import { NotificationDTO } from "../dto/notification.dto";
import { mailerQueue } from "../queues/email.queue";


export const EMAIL_PRODUCER = 'email-producer';

export const addEmailToQueue = async (payload:NotificationDTO) => {
    await mailerQueue.add(EMAIL_PRODUCER, payload);
    console.log(`Email added to queue: ${JSON.stringify(payload)}`);
}