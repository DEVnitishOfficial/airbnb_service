## Steps to setup the starter template

1. Clone the project

```
git clone https://github.com/singhsanket143/Express-Typescript-Starter-Project.git <ProjectName>
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

 

---

# ⚙️ CONFIGURING THE ORM/ODM

---

### ✅ To configure the ORM install the following npm packages:

---

### 1. Install Sequelize (the ORM library)

```bash
npm i sequelize
```

---

### 2. Install MySQL driver library (sequelize uses it internally)

```bash
npm i mysql2
```

---

### 3. Install Sequelize CLI tools

*(generate sequelize related files using sequelize CLI tools)*

```bash
npm install -D sequelize-cli
```

---

## 🚀 To start executing Sequelize, run the following command:

> Starts executing sequelize in the command line.

🛠️ Execute the below command **in `src`** because we have configured all our folder structure inside the `src`.

```bash
npx sequelize-cli init
```

---

### This will create the following folders:

| Folder       | Description                                                                   |
| ------------ | ----------------------------------------------------------------------------- |
| `config`     | 👉 contains the config file, which tells CLI how to connect with the database |
| `models`     | 👉 contains all models for your project                                       |
| `migrations` | 👉 contains all migration files                                               |
| `seeders`    | 👉 contains all seed files                                                    |

---

## 🧩 **Config Explain**

In the `config` folder, there's a `config.json` file. Inside that, there are **three different levels of database connection**:

---

### 1. **Development**

➡️ In order to develop some feature of the application, we have to write some code so writing that code we prefer in development database.

---

### 2. **Test**

➡️ After writing the code, we test that code. So to test that code, we use test database.
➡️ We use test and development database for development and test because if we try it in the production database then if we did any mistake during development, it will impact business.
➡️ Impacting business means loss of money, that we never want.

---

### 3. **Production**

➡️ After successful development and testing of the written code, we push that code in production and then we use the production database.

---

### 4. `"dialect": "mysql"`

➡️ Here through the `dialect` keyword we are saying to Sequelize that we have to connect with the MySQL database, so whatever protocol needs to be enabled in order to connect with the MySQL database, just enable that.
➡️ Because through Sequelize, we not only connect with MySQL, but also with other databases like PostgreSQL, SQLite, Oracle, MariaDB and others, so we have to enable specific protocol.

💡 Or roughly we can say it (`dialect`) **mentions the driver** that from which database we have to connect.

🛠️ **Driver is a raw code written which is used by all ORMs including Sequelize to connect with a particular database.**

---

## 📁 **Explain `models`** (makes interoperable between js/ts and mysql)

Inside the `models` folder, we write the representation (in the form of classes or interface) of the MySQL table in JavaScript/TypeScript.

➡️ Because JavaScript doesn't know the SQL table, so here models will help us to interact with MySQL database by writing the JS-like code.

---

## 🌱 **Explain `seeders`** (change the data inside the table)

Inside the `seeders` folder, we put the **dummy/seed data** of our database.

➡️ So that whenever a new developer comes, they can understand the database or tables easily with that dummy data.

---

## 🏗️ **Explain `migrations`** (change the structure of the database)

Migration folder is used to create different **versions of the database**.

📘 Example:

> Suppose the Airbnb hotel service management system in which initially there were only two columns in a table — `name` and `address`.
> Now, as the product grows and the requirement also grows, we need to add a `STATUS` table as well.
> So, instead of updating the table directly, we **manage the versions** of the database so that we can move from **V1 → V2 → V3** and also can revert back if needed.

---

## 🧹 Deleting all the folders created by Sequelize CLI and organizing inside `src` folder — But Why???

🔐 By default, it creates `config.json` file inside which all the database credentials are present and in the **production-grade application**, it’s a **security compromise**.

✅ That’s why we:

* Delete all the folders
* Create `config.ts` file in `config` folder where now we can **import our DB credentials from `.env`** file.
* Move rest of the folders like `models`, `seeders` and `migrations` inside the `db` folder and **make the folder structure clean**.

---

## ⚙️ Define `.sequelizerc` file

📌 *Put your `.sequelizerc` file inside the `HotelService`* otherwise Sequelize CLI will not be able to read it.

🧠 Inside the `.sequelizerc` file, we write the configuration of Sequelize which is automatically picked up by the Sequelize CLI.

➡️ It allows us to define settings like the **environment**, **configuration file path**, and paths to **migrations**, **seeders**, and **models** folders.

---

## ✅ After creating the files and folders, import all credentials from `.env` file to the `config.ts` file. Now we are ready to move further.

---

# 🚧 Now We Create Migration

🛠️ Run the command below. It will create a **migration file**:

```bash
npx sequelize-cli migration:generate --name create-hote-table
```

---

## 🔍 Understand working of Migration

Migrations have two parts:

| Part     | Purpose                                                                                   |
| -------- | ----------------------------------------------------------------------------------------- |
| **UP**   | Up part of the migration contains the code which makes new changes inside our database after running the migration. Below is that code in up part of migration. |
| **DOWN** |  Down part of the migration contains the code which will revert the changes made by the migration if we want to rollback.Below is that code in up part of migration.  |

---

### 🔼 `UP` Part (Create Table)

```js
async up(queryInterface) {
    await queryInterface.sequelize.query(`
    CREATE TABLE IF NOT EXISTS hotels(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL, 
    address VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );
    `)
}
```

---

### 🔽 `DOWN` Part (Rollback Table)

```js
async down(queryInterface) {
    await queryInterface.sequelize.query(`DROP TABLE IF EXISTS hotels;`);
}
```

---

📛 **Issue Faced:**
By default it supports JS, but we are trying to do it with TS so I got some error.
👉 So, for today we **execute the migration in JS**.

---

## 🚀 Execute the Migration

```bash
npx sequelize-cli db:migrate
```

---

## 🔙 Revert Back the Current Migration

```bash
npx sequelize-cli db:migrate:undo
```

---

# ✅ Benefits of Migrations

---

### 1.

Database migration helps us to **sequentially update our RDBMS**.
➡️ Because RDBMS has **strict schema**, with the help of migration we can manage our database schema or versions with time.
➡️ Also, we can **revert back** to the previous versions using `'DOWN'` part of migration.

---

### 2.

If we feel like our previous version of the database was more stable and useful according to our business requirement, then we can **revert back easily** to that version.

---

# 🛠️ Fixing Error

Initially we have setup the migration in JS but we are writing our backend in TypeScript so all our core logic should be in TypeScript, so now we have to configure our migration in TypeScript.

So, to convert the migration file in TypeScript we have created a **`sequelize.config.js`** file inside the `config` folder. It contains below code:

```js
require('ts-node/register');
const config = require('./db.config.js');
module.exports = config;
```

The above file `sequelize.config.js` is loaded into the `.sequelizerc` file from where the Sequelize CLI reads the configuration of Sequelize ORM.
From here, although we have written our migrations in TypeScript but

**`require('ts-node/register')`** this line converts our TypeScript code into JS code then it will be executed.
In this way both are satisfied:

* we as a programmer write the code in TypeScript
* Sequelize also understands the code in JS

**In short, when we run migration command (npx sequelize-cli db:migrate) to run this command sequelize look into the .sequelizerc and inside here there is config (config: path.resolve("./src/config/sequelize.config.js")) and inside the config(sequelize.config.js) first line is (require('ts-node/register'); which converts on the go our typescript code to js code and then sequelize easily execute js code.)**



---

## 📌 Next Requirement: Adding One More Migration

Suppose a new requirement came that we have to create a **rating** column in our table that will update asynchronously (eventually update, not immediately), but we would like to fetch the rating as soon as possible.

---

### 🛠️ Generate the Migration

Run the below command to create the next migration:

```
npx sequelize-cli migration:generate --name add_rating_hotel_table
```

---

### 🚀 Run Your Migration

```
npx sequelize-cli db:migrate
```

---

### ⏪ Rollback Your Migration

```
npx sequelize-cli db:migrate:undo
```

---

### 🧩 Adding Scripts to `package.json`

In order to run the migrate and rollback multiple times easily, we have added the below script inside the `package.json`:

```json
"migrate": "sequelize-cli db:migrate",
"rollback": "sequelize-cli db:migrate:undo"
```

---

### 💡 Now Whenever We Run:

* `npm run migrate`
* `npm run rollback`

It will run the corresponding migrate and rollback script.
Rollback will not move us to the previous table rather it move us to the previous migration.



Here is your **formatted README** section with better visual structure, using **your exact words** without any modification:

---

## 🧩 Next Step: CREATING DATABASE MODEL AND CONNECTING TO MYSQL DATABASE TABLE

So the next step is to creating the end-to-end API by which people will interact.

For this we will write the **models layer**.

> **So here model will ensure that whatever schema using the database, with respect to the same schema, we will have `class` available in JavaScript/TypeScript, using which we can interact with database indirectly.**

And this class will represent our database in the JavaScript/TypeScript code.

In our JS/TS code we don't want to interact with the SQL tables directly — we interact with classes, functions, and interfaces. That's why here, to interact with the database, we are going to create models by which we can interact with database **indirectly in object-oriented fashion**.

---

### 📝 To write models, created file `hotel.ts` inside the `models`

```ts
class Hotel extends Model<InferAttributes <Hotel>, InferCreationAttributes<Hotel> > {
    declare id: CreationOptional<number>;
    declare name: string;
    declare address: string;
    declare location: string;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare rating: number;
    declare ratingCount: number;
} 
```

Inside the `hotel.ts` above code is written. Explanation of `declare` and `InferAttributes` is written inside the `hotel.ts` and **Notion TypeScript notes**.

---

### ❓ Now here we have a question:

We have defined the `Hotel` TypeScript class by extending the Sequelize `Model`,
but **how our MySQL database will map to the hotel table?**

To map the `Hotel` TypeScript model with our MySQL database table, we have the **`init` function**.
This `init` function takes **two parameters**:

1. ✅ **First parameter** maps the `Hotel` model properties to the database `"hotels"` table columns like (`id`, `name`, `address`, `location`).
2. ✅ **Second parameter** has key `"tablename"` in which we provide our table name as value.
   It also has key `"sequelize"` in which we provide all the configuration of our MySQL database like the `"dialect"`, `"username"`, `"password"`, `"database name"` etc.
3. ✅ To see all these, go inside the `hotel.ts` and `sequelize.ts` file inside the `models` folder.

---

### ✅ Now almost I have solved our problems of database connection and mapping with the `hotel` table in our database successfully.

---

### 🧪 Now we have to insert data into the tables — but how we will do it??

* To insert data inside the table, first go inside the `server.ts` file.
* Check the database connection using `sequelize.authenticate()` method.
* If it is successful, then create a hotel with all defined attributes in `Hotel` model like (`id`, `name`, `address`, ...) using `Hotel.create()` method.

---

### 🚀 **And finally when I run `npm run dev`**,

it creates new entries in the given `hotels` table of MySQL database.

---

### 📖 **To read the data** that we have inserted:

We can use:

```ts
Hotel.findAll()
```



It will return all the created data inside the database.
There are various commands like this we can use.

---

### 🔍 **Fetching Data (Read)**

| Method              | Description                                                |
| ------------------- | ---------------------------------------------------------- |
| `findAll()`         | Returns all entries from the table.                        |
| `findOne()`         | Returns the first record that matches the query condition. |
| `findByPk(id)`      | Finds a record by primary key.                             |
| `findOrCreate()`    | Looks for a record, and creates it if it doesn’t exist.    |
| `findAndCountAll()` | Returns data with the total count — useful for pagination. |

---

### 🆕 **Creating Data (Insert)**

| Method               | Description                                                      |
| -------------------- | ---------------------------------------------------------------- |
| `create(data)`       | Creates a new record in the table.                               |
| `bulkCreate(data[])` | Inserts multiple records in one go (array of objects).           |
| `build(data)`        | Creates an instance but doesn’t save to DB unless you `.save()`. |

---

### ✏️ **Updating Data**

| Method                    | Description                                                     |
| ------------------------- | --------------------------------------------------------------- |
| `update(values, options)` | Updates one or more records matching the condition.             |
| `set()`                   | Updates values on an instance (must call `.save()` after this). |
| `save()`                  | Persists changes made to a built instance or updated values.    |

---

### ❌ **Deleting Data**

| Method               | Description                                             |
| -------------------- | ------------------------------------------------------- |
| `destroy(options)`   | Deletes records matching the condition.                 |
| `instance.destroy()` | Deletes the specific record (used on instance objects). |

---

### 🔁 **Reloading / Refreshing**

| Method      | Description                                 |
| ----------- | ------------------------------------------- |
| `reload()`  | Reloads the instance data from the DB.      |
| `restore()` | Used with soft deletes to restore a record. |

---

### 🧠 **Other Useful Methods**

| Method                   | Description                                           |
| ------------------------ | ----------------------------------------------------- |
| `count()`                | Returns the number of records matching the condition. |
| `increment('field', {})` | Increments a numeric field value.                     |
| `decrement('field', {})` | Decrements a numeric field value.                     |
| `aggregate()`            | Runs SQL aggregate functions like `SUM`, `AVG`, etc.  |
| `max('column') / min()`  | Returns maximum or minimum value from a column.       |

---

---

## ✅ NEXT STEP: WRITING END-TO-END API'S

---

### 🔁 **APPROACH: BOTTOM-UP APPROACH**

> *The bottom-up approach in API writing focuses on building from the ground up, starting with the smallest, most fundamental components and gradually integrating them to create a functional and well-structured API.*

---

### 🗂️ REPOSITORY LAYER

* We're following the **bottom-up approach**, so we begin with the **repository layer**.
* Created a `repository` folder inside `src`, and within it, a `hotel.repository.ts` file where we write all the **DB interactions**.

✅ **Implemented Methods**:

1. `createHotel`
2. `getHotelById`

---

### 🧾 DEFINING DTO

* In `createHotel` → we need hotel data that comes from the **Postman/client (browser)**.

* For **Data Transfer Object**, we define a DTO:

  * Created a `dto` folder.
  * Defined the `hotelData` datatype as `createHotelDto`.
  * Used `createHotelDto` to type `hotelData`.

* Similarly, defined `getHotelById` using `findByPk()` with a **custom error message** via `NotFoundError` from the `utils` (custom error module).

---

### ⚙️ SERVICE LAYER

* Repository layer is now consumed by the **service layer** which handles all **business logic**.

🛠️ Setup:

* Created a `services` folder inside `src`.
* Created `hotel.service.ts` file.

✅ **Implemented Services**:

1. `createHotelService`
2. `getHotelByIdService`

💡 Example of Business Logic (Commented Out):

* Suppose there's a list of blacklisted hotels by address.
* If someone tries to create a hotel using a blacklisted address:

  * It throws a `BadRequestError` and prevents creation.
* This is an example of **pure business logic** — this folder structure supports **separation of concerns** and **Single Responsibility Principle (SRP)**.

---

### 🧭 CONTROLLER

* The **controller layer** now utilizes the service layer.
* Created `hotel.controller.ts` inside the `controllers` folder.

---

### 🧩 ROUTER

* The **controller** is then used by the **router**.
* Created `hotel.router.ts` inside `routers/v1`.

🔁 Routing Setup:

* In `index.router.ts`, initial routing is configured.
* If someone visits:
  `http://localhost:3000/api/v1/hotels`
  → it's redirected to `hotelRouter`
  → on route `'/'` with a **POST** request → it goes to:

  * `createHotelHandler` →
  * `createHotelService` →
  * `createHotel()` →
  * finally to `Hotel` model (where DB schema is defined)
  * then saved to the MySQL database.

---

### ✅ VALIDATION LAYER

* Now we implement the **validation layer**.
* This layer ensures:

  * Incoming data is in the correct format.
  * Prevents future issues in storing/retrieving data.

---

### 🧪 TESTING

* After the validation layer, we **test our APIs**.

🧰 Tools Used:

* Created a **Postman collection** named `"airbnb"`.
* Tested:

  * `createHotel` endpoint
  * `getHotelById` endpoint

📌 Optimization:

* Defined a **Postman variable** to avoid repeating the base URL:
  `http://localhost:3001/api/v1`

---



## 🧠 UNDERSTANDING API FLOW → `createHotel`

## 🔁 REVISION

At first we hit the router in Postman:
`http://localhost:3001/api/v1/hotels` to create a hotel in MySQL database.

---

### 1️⃣ **Request Initiation**

When we hit the above URL, our server accepts that request because we are listening on the same port that is hit by Postman or client.
Then the first request comes to `server.ts` file to:

```ts
app.use(express.json());
```

---

### 2️⃣ **Parsing Request Body**

From Postman we are sending data in `req.body` with:

```
Content-Type → application/json
```

This is parsed by:

```ts
app.use(express.json());
```

and then the request moves further.

---

### 3️⃣ **Attaching Correlation ID**

Then in the same `server.ts` file we have a middleware:

```ts
app.use(attachCorrelationIdMiddleware);
```

In this, on every request we attach a correlation ID to the request headers:

```ts
req.headers['x-correlation-id'] = correlationId
```

This helps to generate a **unique correlation ID for each request** for debugging purposes.

---

### 4️⃣ **Routing the Request**

Next middleware:

```ts
app.use('/api/v1', v1Router);
```

* If our API starts with `/api/v1`, we move them to `v1Router` (inside `index.router.ts`)
* Then, if the request starts with `/hotels`, it is routed to:
  `routers/v1` → `hotel.router.ts`

---

### 5️⃣ **Validation Layer**

After the routing layer, it checks:

* If there's **nothing written after** `/hotels`
* It is a **POST** request

Then the request is sent to the validation layer:

```ts
validateRequestBody(hotelSchema)
```

* Here we have `validateRequestBody` function which takes `hotelSchema`
* `hotelSchema` is a **Zod schema** which validates the required and optional entries in the database that we are going to create

Then we pass this hotel Zod schema to `validateRequestBody`.
This function simply takes:

* The Zod schema
* `req.body`
  → and parses it using the hotel Zod schema.
  If it is successfully parsed, then it calls the `next()` function.

Also, here you can see the same correlation ID because the request is the same.

---

### 6️⃣ **Controller Layer**

After completing the middleware function, the request moves to:

```ts
createHotelHandler
```

We go inside:
`src/controller/hotel.controller.ts`
Here we simply take the `req.body` and call the:

```ts
createHotelService
```

---

### 7️⃣ **Service Layer**

Now we move to:
`/src/services/hotel.service.ts`

Here we write our **business logic**, but for now we're simply creating the hotel.
So we take the `hotelData` and call the:

```ts
createHotel
```

(from repository layer)

---

### 8️⃣ **Repository Layer**

We move to:
`/src/repositories/hotel.repository.ts`

Here in the **repository layer**, we take the `Hotel` model which is the schema representation of our MySQL database in TypeScript — which defines how our database entries look like.

Then, using the **Sequelize `create()` method**, we make entries inside our MySQL database — because we already have `hotelData` that is coming from Postman.

---

✅ After successfully creating the hotel:

* We return from the **repository layer** to the **service layer**
* Then return from **service** to the **controller**
* From **controller**, we return the response back to the **client/Postman**

---

**That's it.**

*Added http-status-codes for enhancing readibility of the code.*


---

## 🧹 Next Step : Soft Deleting Data

Sequelize has a built-in soft deletion feature called `paranoid`. When enabled (`paranoid: true`), deleted records are not removed from the database but are instead marked with a timestamp in a special column. However, we are **not** going to use Sequelize's native `paranoid` functionality, as it may not work consistently across all parts of a real-world application.

---

### 💡 Real-World Insight

In real-world projects, data is rarely **hard deleted**. Instead, it is typically **disabled** or **hidden** from the user interface while still being retained in the database for auditing or recovery. This practice is crucial for data integrity and traceability.

This concept is known as a **tombstone** — meaning the record still exists in the database, but is treated as deleted when it has a timestamp in a specific column.

---

### 🛠️ Our Approach

We’ll implement **manual soft deletion** using a `deleted_at` column in the `hotels` table:

* **By default**: `deleted_at` will be `null`.
* **When soft deleted**: `deleted_at` will contain a `timestamp`, indicating that the record is no longer active.

---

### 🧱 Step 1: Create a Migration

Generate a migration to add the `deleted_at` column:

```bash
npx sequelize-cli migration:generate --name add-deleted-at-to-hotels
```

Then update the migration file:

```ts
module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.addColumn('hotels', 'deleted_at', {
      type: DataTypes.DATE,
      allowNull: true,
      defaultValue: null,
    });
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.removeColumn('hotels', 'deleted_at');
  }
};
```

* `addColumn`: Adds the `deleted_at` column to the `hotels` table.
* `removeColumn`: Removes the column if the migration is rolled back.

---

### 🚀 Step 2: Run the Migration

```bash
npm run migration
```

This command applies the migration and updates the schema by adding a `deleted_at` column in the `hotels` table.

---

### 🏗️ Step 3: Update Codebase

After the database change:

1. Update the TypeScript **Hotel model** to include the `deleted_at` property.
2. Implement `softDeleteHotel` functionality:

   * Repository Layer
   * Service Layer
   * Controller Layer
   * Routing Layer

This flow allows the API to **soft delete** hotels by updating the `deleted_at` field rather than removing the row from the database.

---

Next goal may be to make available room for future 60 days always automatically.

# Now, since we have implemented the migration and model now we will refractore some code but before that we have two question : 

1. Do we need CRUD operation for roomCategory ?
Ans : yes we need, because we the owner can create various type of room cateogory and according to need they can update and delete as well.

2. Do we need CRUD operation for Room ?
Ans : Here answer is both yes and no, for the room creation we do not need of rest network api because for room creation we are going to use the cron job which will automatically create room for next 60 days, but we need a rest api in case of room update and delete.

Now simply i have to implement the CRUD API for room and roomCategory.

# Refactoring of repository layer of hotelService.

If you see the repository layer then we have methods like 
* getHotelById
* updateHotelById
* softDeleteById
* hardDeleteById etc.

These all method we can re-use it because the repository layer nothing have to do with business logic these are just interacting with database so we can make a generalise the repository layer, for this we will use inheritance, class and object.

1. first create a file named base.repository.ts
    visit file for more information and detail that how we have implemented the common base repository and explanation. 

## In previous commit we did room generation synchronously but now we will generate room asynchronusly like we will add job to the redis queue and then processor or worker take the task from that queue and one by one they will execute the task.


## We have generated room asynchronously using the redis(ioredis) and the queue(bullmq) for next 3 months i.e for 90 days, now we want to maintain this 90 days windows that's why here we need cron job expression which will maintain the 90 days window and every day it will generate rooms after 12:00pm for next day.

**So our next goal is to execute the CRON job ---> for this we will use node-cron** 


## 🧭 **Elasticsearch Integration – Hotel Search System**

### 📖 **Overview**

Elasticsearch is a **powerful distributed search and analytics engine** built on top of Apache Lucene.
It allows you to **store, search, and analyze large volumes of data quickly** — even if the data is partially matched, misspelled, or spread across multiple fields.

In this microservice, **Elasticsearch is integrated to provide advanced hotel search functionality** for the Airbnb-like system.
It enables users to search for hotels efficiently using fields like:

* **Hotel Name**
* **Address**
* **Location**

and supports **fuzzy matching**, **partial search**, and **pagination** for better scalability and performance.

---

### 🚀 **Why We Use Elasticsearch**

In a system like Airbnb, users often type incomplete or misspelled names (e.g., “Chnai” instead of “Chennai”).
A traditional SQL-based query (`LIKE '%name%'`) performs poorly with large datasets and cannot handle typos.

Elasticsearch solves this by:

* Tokenizing and indexing text fields (like name, address, location)
* Allowing **fuzzy matching** (handles misspellings)
* Supporting **multi-field relevance search**
* Offering **high performance full-text search**
* Providing **ranking (_score)** to sort the most relevant results

---

### ⚙️ **How It Works – High-Level Flow**

#### 🏨 1. **When a new hotel is created**

1. The HotelService first stores the hotel details in the MySQL database.
2. Then it **triggers a background job** using **BullMQ + Redis Queue**.
3. This job sends the hotel’s data to the **Elasticsearch Worker**.
4. The Worker fetches the full hotel details (including rooms and room categories) from the DB.
5. It transforms this data into an **Elasticsearch-friendly document**.
6. Finally, it indexes (stores) the document in the **Elasticsearch index (`hotels`)**.

```mermaid
graph LR
A[Create Hotel API] --> B[Save in MySQL]
B --> C[Push Job to Redis Queue]
C --> D[Elasticsearch Worker]
D --> E[Index Document in ES Index "hotels"]
```

---

#### 🔍 2. **When a user searches for hotels**

1. The client (e.g., Postman or frontend app) calls `/api/v1/hotel/search?name="Crowne%20Plaza%20Greater`.
2. The `SearchController` passes the search payload (name, address, location) to the `SearchService`.
3. The service constructs an **Elasticsearch query** using:

   * `multi_match` for searching across name and address
   * `match` for location
   * `fuzziness: "AUTO"` to handle typos
   * Pagination using `from` and `size`
4. Elasticsearch returns a ranked list of matched hotels.
5. The response is formatted and sent back to the client.

---

### 🧩 **Core Components**

| Component                           | Description                                                           |
| ----------------------------------- | --------------------------------------------------------------------- |
| **ElasticsearchRepository**         | Handles indexing, deleting, and searching documents in Elasticsearch. |
| **hotelIndexingProcessor (Worker)** | Background worker that listens to Redis Queue and indexes hotel data. |
| **SearchService**                   | Builds search queries and executes them in Elasticsearch.             |
| **transformHotelToESDoc()**         | Converts Sequelize model data into Elasticsearch-friendly structure.  |
| **hotelIndex.queue.ts**             | Defines queue configuration for indexing jobs.                        |
| **hotelIndex.producer.ts**          | Pushes hotel indexing/deletion jobs to Redis.                         |

---

### 🧠 **Important Elasticsearch Concepts Used**

| Concept               | Description                                                              |
| --------------------- | ------------------------------------------------------------------------ |
| **Index**             | Equivalent to a table in SQL. Here we use the `hotels` index.            |
| **Document**          | Equivalent to a row. Each hotel record is one document.                  |
| **Field**             | Equivalent to a column (e.g., `name`, `address`, `location`).            |
| **Analyzer**          | Breaks down text into tokens to make search faster and more flexible.    |
| **Fuzziness**         | Allows typo-tolerant search, e.g., searching "Chenai" matches “Chennai”. |
| **Multi-Match Query** | Searches across multiple fields like name and address simultaneously.    |
| **_score**            | Relevance score assigned by Elasticsearch to rank search results.        |
| **Pagination**        | Uses `from` and `size` parameters to fetch limited results efficiently.  |

---

### 🧱 **Example Query**

```json
{
  "query": {
    "bool": {
      "must": [
        {
          "multi_match": {
            "query": "Chennai Marina",
            "fields": ["name^3", "address^2"],
            "fuzziness": "AUTO"
          }
        }
      ],
      "filter": [
        {
          "match": { "location": "Chennai" }
        }
      ]
    }
  },
  "from": 0,
  "size": 5,
  "sort": [
    { "_score": { "order": "desc" } }
  ]
}
```

---

### 📦 **Example Response**

```json
{
  "success": true,
  "message": "Hotels fetched successfully",
  "data": {
    "total": 11,
    "took": 36,
    "page": 1,
    "size": 5,
    "hotels": [
      {
        "id": "92",
        "name": "Chennai Marina View",
        "address": "Kamaraj Salai, Opp. Marina Beach, Mylapore, Chennai",
        "location": "Chennai"
      }
    ]
  }
}
```

🕐 **`took`:** Time (in ms) Elasticsearch took to execute the search query.

---

### ⚡ **Indexing Flow Summary**

| Step | Action               | Component                              |
| ---- | -------------------- | -------------------------------------- |
| 1    | Hotel created in DB  | `createHotelService()`                 |
| 2    | Add job to Redis     | `hotelIndexProducer.addJob()`          |
| 3    | Worker consumes job  | `hotelIndexingProcessor`               |
| 4    | Transform hotel data | `transformHotelToESDoc()`              |
| 5    | Index in ES          | `ElasticsearchRepository.indexHotel()` |
| 6    | Search hotels        | `SearchService.searchHotels()`         |

---

### 🧰 **Environment Variables**

| Variable             | Description                  | Example                  |
| -------------------- | ---------------------------- | ------------------------ |
| `ES_NODE`            | Elasticsearch Node URL       | `https://localhost:9200` |
| `ES_USERNAME`        | Elasticsearch Username       | `elastic`                |
| `ES_PASSWORD`        | Elasticsearch Password       | `mypassword`             |
| `ES_HOTEL_INDEX`     | Name of index for hotels     | `hotels`                 |
| `ES_BULK_CHUNK_SIZE` | Chunk size for bulk indexing | `500`                    |

---

### ✅ **Best Practices Followed**

* Asynchronous indexing using **Redis + BullMQ** (to avoid blocking HTTP requests)
* **Error handling & retries** in worker
* **Fuzzy & partial search** for flexible results
* **Pagination support** for scalable performance
* **Secure connection** to Elasticsearch with credentials
* **Separation of concerns** (Controller → Service → Repository → Worker)
---


