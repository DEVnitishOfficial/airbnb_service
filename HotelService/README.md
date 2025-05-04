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


Here is your formatted README section, with visually appealing structure and no changes to your original wording:

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

---

Let me know if you'd like to combine this section with the previous formatted README!




