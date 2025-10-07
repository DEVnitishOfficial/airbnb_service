If you make any changes inside the airbnb folder or in future if you want to add any service here then make sure first of all add you particulr service in the main repo like

git add BookingService

Till so far what we have Achieved in HotelService/BookingService/NotificationService

# Hotel Service

1. Configure ORM - sequelize.
    * used mysql2 - driver liberary of mysql
    * install sequelize-cli which generate following file
        → config, model, migration, seeders
        → config.json converted into config.ts
        → define .sequelizerc file
        → after writing model we generate hotel migration with up and down part
        ► written the sql command to create table
        ► After writing the sql command we move to the writing api calls

2. NEXT STEP: WRITING END-TO-END API'S FOR HOTEL SERVICE LIKE
    * written api in bottom-up fahion/approach
    * Repository layer(needed dto layer as well) → Service layer → Controller layer → Router layer → from the router layer we are matching the url and perform the required or defiend action. 

* Then we have written following api
    i. createHotel
    ii. getHotelById
    iii. getAllHotel
    iv. updateHotelById
    v. softDeleteHotelById
    vi. hardDeleteHotelById
    vii. allSoftDeletedHotels
    vii. restoreSoftDeletedHotelById

# Booking Service

1. Configuring ORM : Prisma
     * at first we install the prisma
     * connected to db using mysql connection string.
     * then install prisma client → who interect with mysql db
     * then defined prisma schema
     * finally run the migration with "npx prisma migrate dev --name init"
* Difference between prisma and sequelize
    * in sequelize we define our model(like Hotel) by extending the sequelize Model
    * in prisma we don't need to define the model manually instead
         * when we run "npx prisma migrate" then automatically prisma read .env and schema.prisma file
         * and generates the Prisma client in the node_modules folder, from here whatever types needed we can get easily.

2. Create Booking with Idempotency
    * The idea behind the idempotency were that if a user started transaction and that transaction has not completed yet due to some network issue or whaterver may be various reasons, then doesn't matter how many time they hit the same api we will not allow them to make any changes until a specific time(TTL), if in that given time booking not happen then we will discard the whole transaction/process and user can freshly start booking again.

    * if user has already made successfull booking, still they are trying to make booking with the same userId and hotelId then we will send them a message that they have already created the booking with same userId and hotelId and prevent them from the double booking.
        * Here for the time being i am considering the hotelId as room id because till now we don't have the room id and the start and the end date of a booking, so a user can't book the same room on the same day that they have already booked.

**How we achieved idempotent api with handling concurrency** : 
    * used uuid for idempotencykey generation(generated while creating booking asynchronously)
    * used  redis, ioredis and redlock to acquire lock here(in redlock) we define the resource(on which we have to put lock) and ttl(how much time redis hold that lock on the given resource when ttl completes then lock released then other user can acquire lock on the resource.)
    * Here redis and redlock ensure that no two user can lock the same resource on the same time.
    * So in the above one we have solve concurrency problem(no two user on the same time create booking with the same resource) and make our api idempotent.

**Solving concurrency in confirm Booking(Row level locking in mysql database)**
        **Pessemistic lock, isolation level----> read committed**

        ```ts
export async function confirmBookingService(idempotencyKey: string) {
    const idempotencyKeyData = await getIdemPotencyKey(idempotencyKey);
    if (!idempotencyKeyData) {
        throw new NotFoundError("Idempotency key not found");
    }
    if (idempotencyKeyData.finalized) {
        throw new BadRequestError("Booking already finalized");
    }
    const booking = await confirmBooking(idempotencyKeyData.bookingId);
    await finalizeIdempotencyKey(idempotencyKey);
    return booking;

```

Problem : 

Here in the confirmBooking the main problem were the two or more concurrent confirm booking request, suppose two request came concurrently to our server, now consider these two reqeust as two process that wants to execute in our single core machine,

* So first request process our cpu and waiting for result because it's a asynchronours operation it will not wait for their result and again it will take the second request also it will be processed, suppose second request come fast and due to any network issue first process become delay so the second process proceed and they will confirm the booking and now suppose the first process also resolved and they came with the same idempotency key and trigger the confirmBooking which is already confirmed in the second reqeust, so it's a redundent work that we are doing and it may lead to problem if we are keeping boolean in our db if second request confirm the first one may cancel it.

Solution : 
            So clearly here we can see that the two transactions are reading the same data and trying to modify it, so it's a dirty read problem, to solve this problem mysql provide isolation level of read committed, using which we can put a lock(pessemistic lock) on row level of that idempotency key which it matching form the db.

            With putting pessemistic lock on row level, we will also wrap all the operation in one single transaction like getIdemPotencyKeyWithLock, confirmBooking and finalizeIdempotencyKey these all are the single-single operation we wrap all thse operaiton in one single transaction so that if any of them failed all operation will be failed and then a fresh request will take lock on the same resource this helps in data consistency.

            
```ts
            export async function getIdemPotencyKeyWithLock(tx: Prisma.TransactionClient, key: string) {
    if (!isValidUUID(key)) {
        throw new BadRequestError("Invalid idempotency key format");
    }

    const idempotencyKey: Array<IdempotencyKey> = await tx.$queryRaw(
        Prisma.raw(`SELECT * FROM IdempotencyKey WHERE idemKey = '${key}' FOR UPDATE;`)
    );

    console.log("Idempotency key with lock:", idempotencyKey);

    if (!idempotencyKey || idempotencyKey.length === 0) {
        throw new BadRequestError("Idempotency key not found");
    }

    return idempotencyKey[0];
}
```

# Notification service

# Aggregate rating : 

In this commit i have added the caculation of aggreagte rating of a particular hotel with cron job which in test mode calculate every 30 sec and in production env calculate hotel review hourly.

and solve a bug which took more than 4 hours to match the time basically the mysql by default store time internally in UTC format but we can see them in IST(system) format, so the problem occur when we make query from my server to pickup the recent review which has is_synced false see the below query : 

```go
cutoff := time.Now().UTC()
	log.Printf("ProcessPendingRatings: running cutoff=%s\n", cutoff.Format(time.RFC3339))

rows, err := s.DB.QueryContext(ctx, `
SELECT hotel_id, SUM(rating) AS total_rating, COUNT(*) AS cnt
FROM reviews
WHERE is_synced = FALSE AND created_at <= ?
GROUP BY hotel_id
`, cutoff)
```

I thought if it store it in the local time then make the cutoff call to local so i did cutoff := time.Now().Local() but it's also doesn't work because whenever we make querey our query is interecting with the stored type i.e the UTC internally, but it show us in localtime.

So lastly i decide to make everything in UTC then i converted mysql to the UTC, now it's internally stores in utc and also show us in UTC and form server we make call in UTC so everything works smoothly now.

if you want to check which timeZone use you mysql workbench you can run below command
```sql
SELECT @@global.time_zone, @@session.time_zone;
```
if it shows system means you are using the local time but if it show +00:00	+00:00 then you are using the UTC time.



