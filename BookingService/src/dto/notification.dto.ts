export interface NotificationDTO {
    to: string; // Email address of the recipient
    subject: string; // Subject of the email    
    templateId: string; // ID of the email template to be used
    params: Record<string, any>; // Parameters to be passed to the email template
}