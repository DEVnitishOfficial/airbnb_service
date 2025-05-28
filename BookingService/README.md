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

 Here approach 2 looks good and scalable, but here we can go with anyone so the simplest one we will choose the first approach where first we will create a booking and then idempotency key.

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