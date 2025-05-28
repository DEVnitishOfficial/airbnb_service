import express from 'express';
import { validateRequestBody } from '../../validators';
import { bookingSchema } from '../../validators/booking.validator';
import { confirmBookingHandler, createBookingHandler } from '../../controller/booking.controller';

const bookingRouter = express.Router();

bookingRouter.post('/', validateRequestBody(bookingSchema), createBookingHandler);
bookingRouter.post('/confirm/:idempotencyKey', confirmBookingHandler); 

export default bookingRouter;