import cron from 'node-cron';
import { PrismaClient } from '@prisma/client';
import axios from 'axios';

const prisma = new PrismaClient();
const RELEASE_UNCONFIMED_BOOKINGS_URL = process.env.RELEASE_UNCONFIMED_BOOKINGS_URL || 'http://localhost:3001/api/v1/rooms/release';

 const roomCleanupCron = cron.schedule('* * * * *', async () => {   // runs every minute
  console.log("Running cron job : Release expired bookings...");
  const currentTime = new Date();

  const expiredBookings = await prisma.booking.findMany({
    where: {
      status: 'PENDING',
      expiredAt: {
        lt: currentTime,
      },
    }
  });

  for (const booking of expiredBookings) {
    try {
      // 1. update booking status to EXPIRED in booking table
      await prisma.booking.update({
        where: { id: booking.id },
        data: { 
            status: 'EXPIRED',
            releaseAt: new Date(), 
        },
      });

      // 2. calling HotelService to free the rooms associated with this booking
      await axios.post(RELEASE_UNCONFIMED_BOOKINGS_URL, {
        bookingId: booking.id,
      });

      console.log(`Booking ${booking.id} expired & rooms released.`);
    } catch (error) {
      console.error(`Error releasing booking ${booking.id}:`, error);
    }
  }
});

export function startBookingCleanupCronJob() {
  roomCleanupCron.start();
  console.log("Booking cleanup cron job started.");
}
