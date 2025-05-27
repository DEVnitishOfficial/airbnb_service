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
Now everything is in good state now we have to query the database, in order to query the database we have already installed the prisma client. To handle the query we have to create the instance of the prisma client.

Inside the Prisma create a file ---> client.ts Inside the client.ts created instance of prsma client see below code

import { PrismaClient } from "@prisma/client";

export default new PrismaClient();

* If you remember in sequelize orm we are creating class and extending the model of sequelize suppose if we have create a table named 'Hotel' then we extend the sequelize model and define the property of the Hotel and their respective types like below one

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


So all the above stuff whatever we were doing in the sequelize we don't have to do it in the prisma, here there is something called "Generated types" which generate types for us using the model that we have defined in schema.prisma file, Nothing we have to do manually.

For this use the command "npx prisma generate" in src folder, after executing this command, it reads two files 1. .evn 2. schema.prisma and generated a prisma client in node_modules. And now finally we can start using the prisma client and querying to the database.

So now whenever we need types then prisma client will provides us the types that we have defined in the schema.prisma like if we need the type of "Booking", "BookingStatus"

Now i have created a folder named "repositories" inside this "booking.ts" file.
Then created "services" folder inside this we have booking.service.ts file.

## Create Booking
Now in order to create booking we need an idempotent key so that we can create a unique booking and avoid to duplicate booking.

For this i have to create a model
we can keep the booking id as idempotent key because bookingId is also going to be unique but here we will use the UUID.

Now, when we are going to create a pending booking, on that time we are also going to create a idempotency key in service layer, and this idempotency key will be the nothing but the simple uuid, and this idempotencey key will be used in complete flow.

To generate the uuid, install the npm package 'npm i uuid'

Since we have added the Idempotency table in terms of mysql or model in terms of prisma so we have to create a new migration so i have run the migration using below command

```bash
npx prisma migrate dev --name added_idempotency_key
```