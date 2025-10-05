import { PrismaClient, BookingStatus as PrismaBookingStatus } from '@prisma/client';
import { loadEnv } from '../../config';

// const prisma = new PrismaClient()

const BookingStatus = PrismaBookingStatus;


const bookingSeeds = [
  // Booking 1: Basic Confirmed
  {
    userId: 101,
    hotelId: 1,
    checkInDate: new Date('2025-11-10T14:00:00Z'),
    checkOutDate: new Date('2025-11-15T11:00:00Z'),
    bookingAmount: 125000, // $1250.00 (assuming cents/lowest unit)
    status: BookingStatus.CONFIRMED,
    totalGuests: 2,
    roomCategoryId: 1,
    // Note: createdAt and updatedAt are handled by @default(now()) and @updatedAt
  },
  // Booking 2: Pending, later date
  {
    userId: 102,
    hotelId: 2,
    checkInDate: new Date('2026-01-20T15:00:00Z'),
    checkOutDate: new Date('2026-01-22T10:00:00Z'),
    bookingAmount: 45000,
    status: BookingStatus.PENDING,
    totalGuests: 1,
    roomCategoryId: 2,
  },
  // Booking 3: Recent Completed
  {
    userId: 103,
    hotelId: 3,
    checkInDate: new Date('2025-09-01T16:00:00Z'),
    checkOutDate: new Date('2025-09-07T12:00:00Z'),
    bookingAmount: 210000,
    status: BookingStatus.PENDING,
    totalGuests: 4,
    roomCategoryId: 3,
  },
  // Booking 4: Single night, Cancelled
  {
    userId: 104,
    hotelId: 4,
    checkInDate: new Date('2025-10-05T15:00:00Z'),
    checkOutDate: new Date('2025-10-06T11:00:00Z'),
    bookingAmount: 18000,
    status: BookingStatus.CANCELLED,
    totalGuests: 1,
    roomCategoryId: 1,
  },
  // Booking 5: Long stay, Confirmed
  {
    userId: 101, // Same user as #1
    hotelId: 5,
    checkInDate: new Date('2026-03-01T14:30:00Z'),
    checkOutDate: new Date('2026-03-25T10:00:00Z'),
    bookingAmount: 550000,
    status: BookingStatus.CONFIRMED,
    totalGuests: 3,
    roomCategoryId: 4,
  },
  // Booking 6: Checked In currently (assuming current date is before 2025-10-08)
  {
    userId: 105,
    hotelId: 1, // Same hotel as #1
    checkInDate: new Date('2025-10-02T15:00:00Z'),
    checkOutDate: new Date('2025-10-08T11:00:00Z'),
    bookingAmount: 90000,
    status: BookingStatus.PENDING,
    totalGuests: 2,
    roomCategoryId: 2,
  },
  // Booking 7: Small amount, Confirmed
  {
    userId: 106,
    hotelId: 6,
    checkInDate: new Date('2025-12-01T14:00:00Z'),
    checkOutDate: new Date('2025-12-03T11:00:00Z'),
    bookingAmount: 30000,
    status: BookingStatus.CONFIRMED,
    totalGuests: 1,
    roomCategoryId: 2,
  },
  // Booking 8: Large group, Pending
  {
    userId: 107,
    hotelId: 7,
    checkInDate: new Date('2026-07-04T15:00:00Z'),
    checkOutDate: new Date('2026-07-10T10:00:00Z'),
    bookingAmount: 350000,
    status: BookingStatus.PENDING,
    totalGuests: 6,
    roomCategoryId: 5,
  },
  // Booking 9: Quick weekend getaway, Completed
  {
    userId: 108,
    hotelId: 8,
    checkInDate: new Date('2025-08-16T14:00:00Z'),
    checkOutDate: new Date('2025-08-18T11:00:00Z'),
    bookingAmount: 55000,
    status: BookingStatus.PENDING,
    totalGuests: 2,
    roomCategoryId: 1,
  },
  // Booking 10: Family trip, Confirmed
  {
    userId: 109,
    hotelId: 9,
    checkInDate: new Date('2026-04-10T15:00:00Z'),
    checkOutDate: new Date('2026-04-17T12:00:00Z'),
    bookingAmount: 320000,
    status: BookingStatus.CONFIRMED,
    totalGuests: 5,
    roomCategoryId: 3,
  },
  // Booking 11: Cancelled, long ago
  {
    userId: 110,
    hotelId: 10,
    checkInDate: new Date('2025-03-01T14:00:00Z'),
    checkOutDate: new Date('2025-03-05T11:00:00Z'),
    bookingAmount: 75000,
    status: BookingStatus.CANCELLED,
    totalGuests: 2,
    roomCategoryId: 2,
  },
  // Booking 12: High-roller, Confirmed
  {
    userId: 111,
    hotelId: 11,
    checkInDate: new Date('2026-05-15T15:00:00Z'),
    checkOutDate: new Date('2026-05-22T10:00:00Z'),
    bookingAmount: 780000,
    status: BookingStatus.CONFIRMED,
    totalGuests: 2,
    roomCategoryId: 4,
  },
  // Booking 13: Single guest, low amount, Pending
  {
    userId: 112,
    hotelId: 12,
    checkInDate: new Date('2026-02-14T14:00:00Z'),
    checkOutDate: new Date('2026-02-16T11:00:00Z'),
    bookingAmount: 24000,
    status: BookingStatus.PENDING,
    totalGuests: 1,
    roomCategoryId: 1,
  },
  // Booking 14: Completed in peak season
  {
    userId: 103, // Same user as #3
    hotelId: 13,
    checkInDate: new Date('2025-07-10T15:00:00Z'),
    checkOutDate: new Date('2025-07-15T12:00:00Z'),
    bookingAmount: 195000,
    status: BookingStatus.PENDING,
    totalGuests: 3,
    roomCategoryId: 3,
  },
  // Booking 15: Checked In (Long Stay)
  {
    userId: 113,
    hotelId: 14,
    checkInDate: new Date('2025-09-28T14:00:00Z'),
    checkOutDate: new Date('2025-10-20T11:00:00Z'),
    bookingAmount: 480000,
    status: BookingStatus.PENDING,
    totalGuests: 2,
    roomCategoryId: 4,
  },
  // Booking 16: Confirmed, short notice
  {
    userId: 114,
    hotelId: 15,
    checkInDate: new Date('2025-10-15T15:00:00Z'),
    checkOutDate: new Date('2025-10-18T10:00:00Z'),
    bookingAmount: 72000,
    status: BookingStatus.CONFIRMED,
    totalGuests: 4,
    roomCategoryId: 2,
  },
  // Booking 17: Pending, very far future
  {
    userId: 115,
    hotelId: 16,
    checkInDate: new Date('2027-01-01T14:00:00Z'),
    checkOutDate: new Date('2027-01-08T11:00:00Z'),
    bookingAmount: 290000,
    status: BookingStatus.PENDING,
    totalGuests: 5,
    roomCategoryId: 5,
  },
  // Booking 18: Completed, very low guests
  {
    userId: 116,
    hotelId: 17,
    checkInDate: new Date('2025-06-01T15:00:00Z'),
    checkOutDate: new Date('2025-06-02T10:00:00Z'),
    bookingAmount: 15000,
    status: BookingStatus.PENDING,
    totalGuests: 1,
    roomCategoryId: 1,
  },
  // Booking 19: Cancelled, high amount
  {
    userId: 117,
    hotelId: 18,
    checkInDate: new Date('2026-08-20T14:00:00Z'),
    checkOutDate: new Date('2026-08-27T11:00:00Z'),
    bookingAmount: 600000,
    status: BookingStatus.CANCELLED,
    totalGuests: 3,
    roomCategoryId: 4,
  },
  // Booking 20: Confirmed, family vacation
  {
    userId: 118,
    hotelId: 19,
    checkInDate: new Date('2026-12-20T15:00:00Z'),
    checkOutDate: new Date('2026-12-27T10:00:00Z'),
    bookingAmount: 410000,
    status: BookingStatus.CONFIRMED,
    totalGuests: 4,
    roomCategoryId: 3,
  },
];

loadEnv(); 
const prisma = new PrismaClient();

async function main() {
    console.log('Seeding bookings...');
    for (const bookingData of bookingSeeds) {
        await prisma.booking.create({
            data: {
                userId: bookingData.userId,
                hotelId: bookingData.hotelId,
                checkInDate: bookingData.checkInDate,
                checkOutDate: bookingData.checkOutDate,
                bookingAmount: bookingData.bookingAmount,
                status: bookingData.status,
                totalGuests: bookingData.totalGuests,
                roomCategoryId: bookingData.roomCategoryId
            },
        })
    }
    console.log('Seeding complete! 20 bookings created.');
}

main()

// ... main execution and error handling
