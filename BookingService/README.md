## Steps to setup the starter template

1. Clone the project

```
git clone https://github.com/DEVnitishOfficial/ExpressTypescriptTemplate.git <ProjectName>
```

2. Move in to the folder structure

```
cd <ProjectName>
```

3. Install npm dependencies

```
npm i
```

4. Create a new .env file in the root directory and add the `PORT` env variable


5. Start the express server

```
npm run dev
```

````markdown
# Airbnb Booking Service - Microservices with Prisma ORM and MySQL

This service handles **booking-related functionalities** of the Airbnb system using **Prisma ORM** and **MySQL** as the primary relational database. The goal is to ensure **data consistency**, **transactional safety**, and **isolation**, especially to prevent issues like **double booking** and **concurrency problems**.

---

## 🧠 Key Concepts Covered So Far

- **Understanding Microservices**: What they are and why we use them.
- **Concurrency Handling**: Preventing double bookings and other transactional issues.
- **Database Architecture**:
  - `airbnb_booking_dev`: Dedicated to **booking-related data**.
  - `airbnb_dev_mode`: Handles **hotel and room-related data**.
  - These databases are **separate** for better scalability and deployment flexibility.
  - Direct joins **between the databases are not allowed**.

---

If in case you deleted your db and created new one then run the below command to sync you all table in the new db.

run the command from the prisma file other wise you will see "prisma.schema" file not found.
npx prisma db push


## 🛠️ Setting Up Prisma

### Step 1: Install Prisma

```bash
npm install prisma
````

### Step 2: Initialize Prisma in the `src` Folder

```bash
cd src
npx prisma init
```

### Step 3: Cleanup

* Remove the extra `.env` and `.gitignore` files created by Prisma.

### Step 4: Set up the Database URL

Update the `.env` file with your MySQL connection string:

```env
DATABASE_URL="mysql://<user>:<password>@localhost:3306/airbnb_booking_dev"
```

### Step 5: Install Prisma Client

```bash
npm install @prisma/client
```

> **Note**: `@prisma/client` is the actual library that interacts with your database using the Prisma schema.

---

## 🧩 Prisma Schema Definition

File: `prisma/schema.prisma`

```prisma
model Booking {
  id             Int           @id @default(autoincrement())
  userId         Int
  hotelId        Int
  createdAt      DateTime      @default(now())
  updatedAt      DateTime      @updatedAt
  bookingAmount  Int
  status         BookingStatus
}

enum BookingStatus {
  PENDING
  CONFIRMED
  CANCELLED
}
```

> Prisma now supports **modular schema management**, allowing you to split large schemas into multiple files for better scalability.

---

## 🔄 Running Migrations with Prisma

### Initial Migration

```bash
npx prisma migrate dev --name init
```

### Adding New Migrations

1. Modify the schema (e.g., add a `totalGuest Int` field).
2. Save the file.
3. Run the migration command:

```bash
npx prisma migrate dev --name added_totalGuest_column_to_Booking_table
```

> This will generate a new folder under `prisma/migrations` for each migration.

---

## 🧪 Testing

Once migrations are complete, you can use Prisma Client to interact with your `airbnb_booking_dev` MySQL database.

---

## 📌 Notes

* `airbnb_booking_dev` is **independent** of `airbnb_dev_mode`.
* Prisma schema is defined in a single file by default but can now be split into **multiple files**.
* Designed to support **scalable deployment**, including hosting the booking service on a separate machine.
---

# Prisma Integration and Booking Flow Setup

This project uses **Prisma** as the ORM to interact with the database. Below are the steps and explanations for setting up Prisma and implementing a **Booking** flow using an **idempotency key** to avoid duplicate bookings.

---

## 🛠️ Prisma Client Setup

To query the database using Prisma, we need to create an instance of the Prisma Client.

### Step 1: Create Prisma Client Instance

Create a file `client.ts` inside the `prisma` folder with the following code:

```ts
// prisma/client.ts
import { PrismaClient } from "@prisma/client";

export default new PrismaClient();
```

---

## 🔁 Sequelize vs Prisma

In **Sequelize**, you typically define models by extending the `Model` class. For example:

```ts
class Hotel extends Model<InferAttributes<Hotel>, InferCreationAttributes<Hotel>> {
    declare id: CreationOptional<number>;
    declare name: string;
    declare address: string;
}

Hotel.init({
    id: {
        type: 'integer',
        autoIncrement: true,
        primaryKey: true,
    },
    name: {
        type: 'string',
        allowNull: false,
    },
});
```

### 🆕 In Prisma

You **do not need to manually define model classes**. Prisma generates types automatically based on the `schema.prisma` file.

Run the following command to generate the Prisma client:

```bash
npx prisma generate
```

This reads:

* `.env`
* `schema.prisma`

and generates the Prisma client in the `node_modules` folder.

You can now use types like `Booking`, `BookingStatus` directly from the Prisma client.

---

## 📁 Folder Structure

We have created the following structure:

```
src/
├── prisma/
│   └── client.ts
├── repositories/
│   └── booking.ts
├── services/
│   └── booking.service.ts
```

---

## 🧾 Create Booking with Idempotency

To avoid duplicate bookings, we use an **idempotency key**, which ensures each booking request is uniquely handled.

### 🔑 What is an Idempotency Key?

* A unique key (UUID) generated when a booking is initiated.
* Used throughout the complete booking flow to ensure the request is not duplicated.

### 🧬 How to Generate UUID

Install the `uuid` package:

```bash
npm install uuid
```

Generate the UUID inside the service layer when creating a pending booking.

---

## 🧱 Add Idempotency Table (Model)

Since we added a new model `Idempotency` in Prisma, we need to create a new migration.

Run the following command:

```bash
npx prisma migrate dev --name added_idempotency_key
```

This will:

* Apply the changes in `schema.prisma`
* Update the database schema
* Generate necessary client updates
---

## Setting up Idempotency key based API

In order to generate the idempotency key we need some utility function.

So in the utils--->Helpers folder created a file named "generateIdempotencyKey" which will generate a unique string, we keep it seperate because may be in future if we want to change the implementation then this function will be helpful.

Here we can observe on things that the 'idempotencyKey' model is very similar to the booking model so all the repository function for idempotencyKey can implement in the booking repository. 

In the idempotencyKey model we have added a "finalized" property which will be a boolean and by default it's false, purpose of this finalized property is to, confirming the value of *finalized* is true, if it is true it means idempotencyKey completed their full lifecycle, and a successful booking has been created. if the same idempotency key came for booking again then we don't have to process the query or create a new booking, we can just return the existing booking.

**Ways to create idempotencyKey, booking and then connecting it**
1. Approach One : 
 First approach is to create "Booking" and then create "idempotencyKey", this will be sequential and will take time, because here first db call will create booking and second db call will create idempotencyKey and inside the idempotency key booking instance will be added. Here if the booking is created then only we can create the idempotency key.

 2. Approach Two : 
 In the second approach first we will create a *booking* which will be in the pending state and parallely we create a instance of idempotencyKey, currently we will have no booking and in the seperate request we will add booking instance to the idempotency key which will confirm the booking. In this approach we will get idempotency key earlier and if we want we can make update request asynchronous and send back response to the user, in this way user will get fast or quick response.Here one problem may occur like while we are doing the asynchronous update it may fail, so in our backend system we can re-try to update till 2-3 attemtp.

 Here approach 2 looks good so we will go with this one i.e on the time of create booking(pending state) we will generate the idempotency key asynchronously and in later request we will confirm the booking as well.

 **Working with two tables(making relationship b/w tables)**

Since we have two model/table now so to create a booking we need booking and idempotency key both, so we have to create a relationship between these two.

To make the relationships beteen the models we have to change something from prisma.schema

Previously we were doing one-to-one relationship still we will do one-to-one relationship but here i want to make the booking table clean, so we will establish this relationship from idempotency table

since we make changes in our model(schem.prisma file) so we have to migrate database by following command : 
npx prisma migrate dev --name Added_relationship_for_booking_to_idmpotency_key

Now whenever we have to create a booking then we will use BookingCreateInput function it has ----> userId: number;
          hotelId: number;
          createdAt?: Date | string;
          updatedAt?: Date | string;
          bookingAmount: number;
          status?: $Enums.BookingStatus;
          totalGuests: number;
          idempotencyKey?:Prisma.IdempotencyKeyCreateNestedOneWithoutBookingInput


---

# ✅ Solving Concurrency Problem in `confirmBookingService`

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
}
```

---

## 🧠 Problem

The problem lies in the above function. Suppose a user created their booking and now wants to confirm it. If, by mistake or in a frustrating situation (e.g. application not responding due to network issues), the user sends **two requests very quickly (in milliseconds)**, they almost run in parallel:

* First request gets the idempotency key.
* Context switches to second request, which also gets the same key and **finalizes** the booking.
* Context switches back to the first request, which now tries to finalize the booking **again**.

This causes unnecessary or duplicate operations, since the booking was already confirmed by the second request.

---

## 🛠️ Solution

To prevent this:

1. **Wrap the entire `confirmBookingService` in a single transaction** – for rollback if any operation fails.
2. **Put a pessimistic lock (`SELECT ... FOR UPDATE`)** on the `getIdemPotencyKey`, so only one request proceeds, and the rest are blocked until the first completes.

---

## ❓ Questions

### 1. How to wrap all operations in one transaction?

Since we are using **Prisma ORM**, Prisma provides **interactive transactions** where we pass an async callback into `$transaction`.

#### What is `$transaction`?

* `$transaction([])` is an API provided by Prisma.
* It allows us to **run multiple operations as a single atomic operation** – if any step fails, the entire transaction is rolled back.

#### Notes:

* Inside `$transaction`, we can wrap all related operations.
* The callback receives a parameter `tx`, which is a scoped instance of `PrismaClient`.
* Use `tx` just like `prismaClient`, but all operations will be transactional.
* **If any step fails, everything is rolled back.**

### ✅ Code Example – Wrapping in Transaction

```ts
export async function confirmBookingService(idempotencyKey: string) {
    return await PrismaClient.$transaction(async (tx) => {
        const idempotencyKeyData = await getIdemPotencyKeyWithLock(tx, idempotencyKey);
        if (!idempotencyKeyData) {
            throw new NotFoundError("Idempotency key not found");
        }
        if (idempotencyKeyData.finalized) {
            throw new BadRequestError("Booking already finalized");
        }
        const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
        await finalizeIdempotencyKey(tx, idempotencyKey);
        return booking;
    });
}
```

---

### 2. How to put lock on the `idempotencyKey`?

Since we’ve wrapped our operations in a transaction and have the `tx` instance, we can use SQL's **pessimistic locking** with `SELECT ... FOR UPDATE`.

#### Implementation:

* Use raw SQL with Prisma's `tx.$queryRaw`.
* Validate the key to ensure it is in UUID format.
* Use **parameterized queries** to avoid SQL injection.
* The row matching the `idempotencyKey` will be locked for that transaction.

### ✅ Code Example – Locking on Idempotency Key

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

---

> ✅ This solution ensures **only one request proceeds** when multiple parallel requests are made. All other requests are either blocked or rejected once the booking is already confirmed.

---

# 🔒 Solving the Concurrency Booking Problem

When two different users try to book the **same hotel at the same time**, we need to prevent **race conditions** that could lead to **double bookings**.

To handle this concurrency issue, we use a **distributed locking mechanism** with **Redis** and **Redlock**.

---

## ✅ Problem Statement

In a microservices environment (like hotel booking), multiple instances of the same service might receive concurrent requests. To ensure **only one user** can book a hotel room at any given time, we use a **distributed lock** on the booking resource (hotel ID).

---

## 🚀 Technologies Used

* [`ioredis`](https://github.com/luin/ioredis) – Redis client for Node.js
* [`redlock`](https://github.com/mike-marcacci/node-redlock) – Implements Redlock algorithm for distributed locks

---

## 📦 Installation

```bash
npm install ioredis
npm install redlock
```

---

## ⚙️ Redis Configuration (`config/redis.config.ts`)

We create a Redis client and a Redlock instance, then export both to use across the application.

```ts
import IORedis from 'ioredis';
import Redlock from 'redlock';
import { serverConfig } from '.';

export const redisClient = new IORedis(serverConfig.REDIS_SERVER_URL);

export const redlock = new Redlock([redisClient], {
    driftFactor: 0.01,        // Compensation for Redis TTL drift
    retryCount: 10,           // Retry 10 times before giving up
    retryDelay: 200,          // Wait 200ms between retries
    retryJitter: 200,         // Add randomness to retry to avoid stampede
});
```

---

## 🔐 Locking Implementation in Booking Service (`booking.service.ts`)

We use `redlock.acquire` to put a lock on a hotel using its ID, ensuring only one request can proceed with the booking process at a time.

```ts
export async function createBookingService(createBookingDTO: CreateBookingDTO) {
    const ttl = serverConfig.LOCK_TTL;
    const bookingResource = `booking:${createBookingDTO.hotelId}`;

    try {
        // Acquire a lock on the hotel to prevent concurrent booking
        await redlock.acquire([bookingResource], ttl);

        // Proceed with booking if lock is acquired
        const booking = await createBooking({
            userId: createBookingDTO.userId,
            hotelId: createBookingDTO.hotelId,
            bookingAmount: createBookingDTO.bookingAmount,
            totalGuests: createBookingDTO.totalGuests,
        });

        // Generate idempotency key to prevent duplicate operations
        const idempotencyKey = generateIdempotencyKey();
        await createIdempotencyKey(idempotencyKey, booking.id);

        return {
            bookingId: booking.id,
            idempotencyKey: idempotencyKey
        };

    } catch (error) {
        throw new internalServerError("Failed to acquire lock for booking resource");
    }
}
```

---

## ⏱️ How TTL Works

* **TTL (Time to Live)** is passed as the second argument to `redlock.acquire`.
* It defines how long the lock is held.
* If the booking process doesn't complete within the TTL window, the lock **automatically expires**, allowing others to try again.

---

## 🎯 Result

This approach ensures:

* Only one user can book a specific hotel at a time.
* Other requests will retry or fail gracefully.
* Prevents race conditions in a distributed system.

---

# 📬 Notification Service Setup

In the **NotificationService**, the **worker is ready to consume/process jobs**. From the **BookingService**, we need to add jobs to the **Redis queue** using the `bullmq` `Queue`. It’s crucial that the **queue name** matches exactly, because the **worker listens to a specific queue name** — if the names differ, the worker will not process the job.

---

## 📁 Folder Structure Overview

### 1. `queues/`

Contains the `mailerQueue` instance created using `bullmq`'s `Queue` class.

### 2. `producer/`

Uses `mailerQueue` to add jobs to the `"email-producer"` queue with the appropriate payload.

### 3. `dto/`

Defines the shape/type of the incoming request payload using a `NotificationDTO`.

### 4. `server.ts` (for testing)

Used to test the setup by sending a sample job to the queue.

---

## 🔄 How the Microservices Work Together (BookingService ↔ NotificationService)

### ✅ Redis as the Communication Bridge

* The **BookingService** acts as the **producer**. It schedules a job using `bullmq.Queue` into the **Redis queue** with the queue name `"email-producer"`.

* The **NotificationService** acts as the **consumer/worker**. It listens to the Redis queue named `"email-producer"` and processes jobs as they arrive.

### 🔁 Event Flow

1. **BookingService** pushes a job to Redis:

   ```ts
   mailerQueue.add("email-producer", payload)
   ```

2. **NotificationService's** worker is continuously polling Redis:

   ```ts
   const worker = new Worker("email-producer", async (job) => { ... })
   ```

3. When a job is found in the queue:

   * It is fetched by the worker
   * Processed based on the job payload

### ☑️ Why This Works

* Both services connect to the **same Redis instance** (same host and port)
* Both use the **same queue name** (`"email-producer"`)
* Therefore, they are effectively **communicating via Redis**, even though they are **separate microservices**

> 🔄 This model is **loosely coupled and scalable**, making it a reliable choice for microservice communication.

---


# Now we are going to connect everything : 

* In the booking service i have changed booking model checkInDate checkOutDate and roomCategoryId
* And in the HotelService room table we have already the bookingId which connect booking to rooms

* Now before creating the booking, we have to check is the rooms are available for particular roomCategoryId within certain checkIn and checkout date range or not, how to check it?

Ans : if we have checkInDate and checkOutDate but don't have the bookingId in the rooms table then that particular room is available and we can create booking for that roomCategoryId with specified date range according to requirement.

    here booking service will make an api call to the hotelService to see if any rooms are available for a date range or not, and then according to the return set of value we will create the bookings.

# Till so far we don't exposed any api to find the roomCategoryId and dateRange so in the HotelService i will work to expose that api
