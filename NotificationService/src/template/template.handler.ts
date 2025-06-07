import fs from 'fs/promises'; // fs/promise --> reading file asynchronously
import path from 'path'; // solve the path seperator problem(windows : '\' and linux/macos:'/')
import Handlebars from 'handlebars';
import { internalServerError } from '../utils/errors/app.error';


export async function renderMailTemplate(templateId:string, params: Record<string,any>): Promise<string> {
  const templatePath = path.join(__dirname,'mailer',`${templateId}.hbs`); 
  
  try{
    const content = await fs.readFile(templatePath,'utf-8');
    const finalTemplate = Handlebars.compile(content);
    return finalTemplate(params);
  }catch(error){
    throw new internalServerError(`Template not found ${templateId}`)
  }
}

// __dirname provide the absolute path of the current file like here path of templatePath is
  //  D:\INTERVIEW PREPERATION\SANKET_BACKEND\BACKENT_PROJECT\Airbnb\NotificationService\src\template\mailer\welcome.hbs
  // only __dirname provide till : D:\INTERVIEW PREPERATION\SANKET_BACKEND\BACKENT_PROJECT\Airbnb\NotificationService\src\template i.e wherever you are writing the code till that folder.