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










