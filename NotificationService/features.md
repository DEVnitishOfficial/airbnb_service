# Core Features of NotificationService

## 1. Send notification on provided email address

* In NotificationService there are three key role players : 

(i) Producer : provide the payload which include to, subject, templateId, and params

(ii) Redis : store payload using bullmq

(iii) processors : using worker take those payload from redis and send email

### How it's started working 

- starts from server.ts> addEmailToQueue(payload)

* At first we have to create a job and pass payload to that job here the job is to send the email notification to the end user, and inside that payload we send the userInfo, like "to", "subject", "templateId", and params.

* we register a queue in redis named MAILER_QUEUE inside this queue we have added a job name EMAIL_PRODUCER, under this name we put our payload here the work of producer is ended.

* The similir kind of structure is available in the bookingService as well where there is  (queues > email.queue.ts) in which there is name of the queue and redis-connection also there is producers > email.producer.ts which put paylod under job name EMAIL_PRODUCER

* Here in our NotificationService there is setupEmailWorker function inside this we have implementd the email processing task using the bullmq, here our worker trying to find the job under the queue named MAILER_QUEUE and job name EMAIL_PRODUCER here when they found the payload then extract each info and process the mail info and send to the end user.

* Here for email template we have used handlebars and nodemailer for sending email.

* Handlebars is a simple templating language.




