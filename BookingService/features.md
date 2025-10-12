
How Booking Creation and confirmation works(core business logic)

1. dto : according to dto took payload from the client.

```ts
export type CreateBookingDTO = {
    userId: number
    hotelId: number;
    bookingAmount: number;
    totalGuests: number;
    checkInDate: Date;
    checkOutDate: Date;
    roomCategoryId: number;
}

{
    "userId":911,
    "hotelId":12,
    "bookingAmount":800,
    "totalGuests":5,
    "roomCategoryId": 15,
    "checkInDate": "2026-01-28",
    "checkOutDate": "2026-01-28"  
}
```
2. check available rooms by making an external api calls to the hotelService getAvailableRooms() method

3. if rooms available on that checkout range then **lock the bookingResource(room id)** using using the redis and redlock with ttl.

* Redis : Redis is an in-memory data structure store.
* RedLock : redlock is an algorithem, used to lock the particular resources in a distributed system on multiple independent redis instances.

4. ttl is set for 60sec if in this duration, booking is created the update the db other wise lock will be release others can book or use that resources.

5. When booking is created that stays in pending state, if in this situation any other user come and they try to create the same booking on same date range they will be able to create the booking since till so far booking is not confirmed so it's possible, if booking is confirmed then no-one can checkout that particular date range of a particular hotel roomId in specific date range.

6. In order to confirm the booking user give back that idempotency key and using that idempotency key we confirm the booking.

   
   ### Why you create an ItempotencyKey, what are it's uses here?
    
    * if any operation called as idempotent it means if we do the same operation multiple time then the outcome will remain same.

   Ans : To avoid the double bookings or double charges to the client.

   * Simple flow of booking
            i. First user choose the hotel, room_category, no of guest, checkIn & checkout date etc.

            ii. Once they fill all detail then before payment details we create a booking in pending state and send them a unique booking id in our case uuid(we call it idempotency key)

            iii. When user is filling the payment details and confirm the checkout, then client send that unique booking uuid using that we update the status to confirm and marked the particular date range roomId as booked. 


    **IdempotencyKey usage** ---->>>> Now suppose the client click confirms button or send request through postman then we get that unique uuid bookingId that we say idempotencyKey from the client and before updating the status in db first we check in our *idempotencykey* table  if the key exist and finalise column is also true then we will reject that booking saying, booking has already been finalized.
                In this way no matter client how many times send request if the *finalized* column in the *IdempotencyKey* table is true it will be rejected.


    **Uses of Transaction While confirming booking** ------>>>> while confirming the booking we do three db operation in single transaction : 
    (i). check the *finalized* column from the *IdempotencyKey* table
    (ii). if finalized column founds false then then took the bookingId from the idempotencyKeyData(result of the first query) and on the basis of that booking id we confirm the status of booking from PENDING to CONFIRM in *booking* table
    (iii). After updating the status of booking inside the *booking* table we call back to the *IdempotencyKey* table and update the finalize property to true. So now if any req with same idemKey will come that will be rejected.

    *If any of these operation failed the complete transaction will be rollback and no changes will be applied.


    **What are the database and table in BookinService**

    Here our DB name is : *airbnb_booking_dev*
    And here we have two tables with following detail

    1. IdempotencyKey table
    id                Int              @id @default(autoincrement())
    idemKey           String           @unique
    finalized         Boolean          @default(false) 
    bookingId         Int              @unique 
    createdAt         DateTime         @default(now())
    updatedAt         DateTime         @updatedAt

    2. Booking table 
    id                Int   @id         @default(autoincrement())
    userId            Int
    hotelId           Int
    checkInDate       DateTime
    checkOutDate      DateTime
    bookingAmount     Int
    roomCategoryId    Int
    status            BookingStatus     @default(PENDING)
    totalGuests       Int
    createdAt         DateTime          @default(now())
    updatedAt         DateTime          @updatedAt

