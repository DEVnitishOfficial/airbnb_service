
# How Booking Creation and confirmation works(core business logic)

## working of bookingCreation
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

* await redlock.acquire([bookingResource], ttl); ----handle concurrency
* SET booking:101 "abc123" NX PX 5000

* SET → create or update a key
* booking:101 → represents the room or resource you want to lock
* "random-unique-value" → the unique ID of the current lock owner (used to unlock safely later)
* NX → only set if the key doesn’t already exist (so no overwrite)
* PX 5000 → expire after 5000 milliseconds (TTL = 5 seconds)

4. ttl is set for 60sec if in this duration, booking is created the update the db other wise lock will be release others can book or use that resources.

5. when booking is successfully is created we have the bookingId, then we generate the idempotency key, and call the repository layer since we have idempotencykey table exist so we update our idempotencyKey table with newly created idempotency key and bookingId.

6. After this we set the expiry time of the newly created booking to 10 minutes if within 10 mins booking is not confirmed then the booking is expired i.e the bookingId will be removed from the room table.

7. After updating expiry we make an api call to the hotelService and update the bookingId into the room so that, that particular room will be unavailable for 10mins becuase that room has been booked by someone, if withing 10 mins not confirm again i will remove that bookingId and room will be available for others.

## working Booking Confirmation : 

8. In order to confirm the booking we have a method named confirmBookingService which takes two parameter (i)idempotencyKey (ii)currentLoggedInUserId

9. start prismaClient transaction, since we have idempotencyKey table and client already providing the idempotencyKey which they got during the booking creation, now i have to look into that idempotencyKey table and see if the booking is alredy confirmed or not and accordingly we have to return res, but here is catch.

10. when we have to read the staus of the particular idempotency key before updating their status, so if we found the "finalized" property is false it means the booking is unconfirmed and we have to confirm the booking.

11. Before moving further first of all we have to put a **row level lock(using select.....for update)** so that no other transaction can modify the current transaction until or unless this txn complete or rollback.

12. After putting lock we read the current booking status from the given idempotency key if booking is not already confirmed, then we check is the loggedIn user and the user who is created this booking is same or not, if both user is same then we call

13. finalizeIdempotencyKey repository layer with idemKey, using which it identify the row and update the finalize column to true, i.e booking is confirmed.
   
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

## Implementd corner case like the user who has created booking is the same user who is confirming the bookings.

**How we have implemented**

* To identify the user who is currently confirming the booking is the same user who have created the booking then any how i have to extract the userId from the idempotency key.

* if you see your booking table then there is a connection between the booking table and the idempotency table, in the idempotency table there is bookingId and inside the bookingId there is userId so joining the idempotency table and the booking table we can easily extract out the userId.

* Now we have the userId extracted from the idempotencyKey, and we can extract the current user userId form the current login token and by comparing both if both are same it means the loggedIn user and the user who created this booking is same, so allow the booking confirmation. 

## Other approach to solve the double booking : 

* In our current implementation when booking is successfully created we generate the idempotencyKey and insert it into the idempotencyKey table, since before booking confirmation we are inserting the idempotency key so when we have to confirm the booking on that time we required acquire row level lock because may be same req come with same idem key since we have row level lock so the second request will wait until our current txn completes, and once completed then on the same txn we are updating the finalized propety in the idempotencyKey table, so second txn will be rejected because first txns already confirm the booking.


* The second approach can be when we create the booking, after this we create the idempotency key and do all the rest operation and instantly we don't have to update the idempotencyKey table rather send the bookingId and the idem key to the client and when client try to confirm the booking on that time enter the bookingId and idemKey into the idempotency key and in the same txns also update the finalized property so whenever second request will come with the same idem key then that will be automatically rejected because of the uniqueness of the idem key.

```js
if (booking.finalized) return "Already confirmed";
else {
  markBookingAsConfirmed();
  chargePayment();
}
```
**solved concurrency(multiple user try to book same resource) ---> using redis and redlock**

**solved double charge or booking(same user try to confirm multiple time) ----> using idem key and row level exclusive lock.**