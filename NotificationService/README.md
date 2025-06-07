# Designing Booking service

Theory of booking service already covered at first revise this one from here
[Notion Notes](https://www.notion.so/Notification-Designing-of-bookingService-206ce7b46932809f8468c274cf2ae3cf)

Till so far we have developed the HotelService and BookingService and now we are going to build the Notification service for this to start we need :

**ioredis**:
 npm i ioredis 
 // It's a client who will interect will the redis-server.

**bullmq**: 
npm i bullmq 
// BullMQ is a NodeJS library for creating background jobs and sending messages using queues

**nodemailer** : 
npm i nodemailer 
// it's a Node.js module used for sending emails from server-side applications

**handlebars**: 
npm i handlebars 
// handlebars is a templating language/engine used to generate template.

# Setting redis configuration :

```ts
export function connectToRedis() {
  try {
    const redisConfig = {
      host: serverConfig.REDIS_HOST,
      port: serverConfig.REDIS_PORT,
    };
    const connection = new Redis(redisConfig);
    return connection;
  } catch (error) {
    console.error("Error connecting to Redis:", error);
    throw error;
  }
}
```

Initially we tried to setup the redis configuration like the above one but in the above one there is some problem which can cause the performance issue.

# PROBLEM :

if you see carefully your current code then if a user will call the connectToRedis() function then each time it will create a new connection, suppose if you are using or calling 100 times connectToRedis() function throught your codebase then it will create 100 different Redis instances which can cause below problem :

- it will open multiple tcp connection between the nodejs and redis which further required memory buffer to read or write the data, and that tcp connection may still idle which exhaust the machine resource.

- Redis provide certain number of concurrent clients connection(eg.10K) if we create new connection on every every task/job it may exhaust that connection, resulting it refuse to new connection.

- Race condition problem : When multiple connections perform operations that depend on shared state, but they don’t coordinate — this leads to unexpected behavior, like in publisher/subscriber model suppose If different Redis instances are used for subscribing vs publishing, messages might not be received or process.

- can face problem in caching suppose One connection sets a key with a TTL, and another connection tries to access or delete it before the TTL ends.

- When many Redis connections are spread throughout the codebase, then it will be Hard to trace which connection failed, what caused a spike in Redis traffic etc.

# SOLUTION :

So considering the above problem we have decided to use the singelton function using the clouser concept in which we will write our logic in such a way so that the redis connection is created only once and whenever we need the Redis connection we use the same connection insted of creating the new one on each time, below is the implementation of that code.

```ts
export function connectToRedis() {
  try {
    let connection: Redis;
    const redisConfig = {
      host: serverConfig.REDIS_HOST,
      port: serverConfig.REDIS_PORT,
      maxRetriesPerRequest: null, // Disable automatic reconnection
    };

    return () => {
      if (!connection) {
        connection = new Redis(redisConfig);
        return connection;
      }
      return connection;
    };
  } catch (error) {
    console.error("Error connecting to Redis:", error);
    throw error;
  }
}

export const getRedisConnObject = connectToRedis();
```

The above code is creating only one instance of redis and it is used throughout the component life cycle.

**Queue, Producer, Processer(worker), working**

_Queue_
After writing the singelton function i have used getRedisConnObject() inside the mailer.queue.ts where we have created a Queue using bullmq node liberary, so in order to create a Queue we need the queue name and the a connection string which nothing but the redis host and port, where redis is running, so in the connection we provide getRedisConnObject() which is nothing but the instance o redis.

```ts
import { Queue } from "bullmq";
import { getRedisConnObject } from "../config/redis.config";

export const MAILER_QUEUE = "mailer-queue";

export const mailerQueue = new Queue(MAILER_QUEUE, {
  connection: getRedisConnObject(),
});
```

_Producer/Publisher_
Since we have created a Queue named mailerQueue now in this queue we have to add the task, for this we have a method named "add" and it takes three parameter 1. name of the job and 2.payload(data on which worker will work) 3. delay(waiting time to execute the job).Below is the implenetation. Using these options job will be added to the redis queue with proper queue and job name.

```ts
import { NotificationDTO } from "../dto/notification.dto";
import { mailerQueue } from "../queues/mailer.queue";

export const EMAIL_PRODUCER = "email-producer";

export const addEmailToQueue = async (payload: NotificationDTO) => {
  await mailerQueue.add(EMAIL_PRODUCER, payload);
  console.log(`Email added to queue: ${JSON.stringify(payload)}`);
};
```

_Subscriber/processors/workers_
Now when job is added to the redis queue, then processors/workers come into picture, bullmq automatically fetch jobs from the redis queue and make available for ther worker in the form of "job" keyword, this job has all the property available that has been defiend in the NotificatonDTO.
Below is the implementation :

```ts
export const setupEmailWorker = () => {
  const emailProcessor = new Worker<NotificationDTO>(
    MAILER_QUEUE,
    async (job: Job) => {
      if (job.name !== EMAIL_PRODUCER) {
        throw new Error(`Invalid job name: ${job.name}`);
      }

      //call service layer from here to process the email
    },
    {
      connection: getRedisConnObject(),
    }
  );

  emailProcessor.on("failed", () => {
    console.error("Email processing failed");
  });

  emailProcessor.on("completed", (job) => {
    console.log("Email processing completed successfully for job:", job.name);
  });
};
```
Here worker takes three parameter 1. Name of the queue from where it will pull the job 2. a callback function in which job will be executed and 3. connection of the redis.

# Testing
I have tested this in my server.ts file with sample job like below one :
```ts
const sampleNotification : NotificationDTO = {
        to: "sample",
        subject: "Sample Subject",
        templateId: "sample-template",
        params: {
            name: "Nitish official",
            orderId: "12345",
            message: "This is a sample notification message."

        }
    };
    addEmailToQueue(sampleNotification)
```

# Setting up the email
Till so far our setup is that jobs are produces by the producer from the BookingService and that job is processed in NotificationService by the processor, also till now we are sending the payload from the server file manually but now we have to setup a proper email template and required field according the the NotificationDTO like here we have 
```ts
export interface NotificationDTO {
    to: string; // Email address of the recipient
    subject: string; // Subject of the email    
    templateId: string; // ID of the email template to be used
    params: Record<string, any>; // Parameters to be passed to the email template
    // here Record<string, any> represents key will be string and value may be of any type like number, string, boolean etc.
}
```
In the above one there is templateId which we have to create first, so to create template we can use any templeting engine language like ejs/hadlebars so for simplicity here we will use handlebars.

## Handlebars : 
* Created a folder template inside that we have mailer folder and template.hander.ts, inside the mailer folder we will write all template like for now i have written welcom.hbs file similir to this one there may be different template file.

* To read all these template file we have written code in template.handler.ts file,you can see here how asynchronous file readings are happening.


# Configuring NodeMailer:
Since we have already create the template that i have to sent by email now we need to setup the nodemailer so that we can send mail to the user.

* inside config>mailer.config.ts i have configured the mail transpoter using nodemailer.createTransport and export it so that we can use it in other file where it is needed, visit for more detail implementation.

* Used transpoter in services>mailer.service.ts with config option like to,from subject,body details, and exported sendEmail function from the service layer.

* Then used sendEmail inside the processors>email.producer.ts file with necessery payload and mesage content that we want to send the end user like mail of the user sent using payload.to, subject : payload.subject and emailContent as the body of the mail.

* And finally from the server.ts file called addEmailToQueue and pass the required payload and it successfully sent email to the given email address.



