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

 ## CONFIGURING THE ORM/ODM

 To configure the orm install the following npm packages

At first install the sequelize(the orm liberary)

```
npm i sequelize
```

Then install the mysql driver liberary(sequelize uses internally)
 ```
 npm i mysql2
 ```

 Also install the sequelize cli tools(generate sequelize related files usign sequelize cli tools)

```
npm install -D sequelize-cli
```
## **To start executing the sequelize we have to execute the following command**

starts executing sequelize in command line.

execute the below command in src because we have configured all our folder stracture inside th src
```
npx sequelize-cli init
```

This will create the following folders

config --->> contains the config file, which tells CLI how to connect with the database

models --->> contains all models for your project

migrations --->> contains all migration files 

seeders  --->> contains all seed files

**Config Explain**

In the config folder ther is a config.json file, inside that there are three different level of database connection, 

1.Developement --> In order to develop some feature of application, we have to write some code so writing that code we prefer in develpement database.

2.Test --> Now after writing that code we have test that code so to test that code we will use test database, we are using test and developement database for developement and test because if we try it in the production database then if we did any misteak during the developement then it will impact business, impacting business means loss of money, that we never want.

3.Production --> After successful developement and testing of the written code then we push that code in production and then we use the production database.

4. "dialect": "mysql" ---> here through the dialect keyword saying to the sequelize that i have to connect with the mysql database so whatever protocol need to enable in order to connect with the mysql database just enable that. Because through the sequlize we not only connect with the mysql alos we can connect with other database like PostgreSQL, SQLite,oracle,mariadb and others so we have to enable specific protocol.

Or roughly we can say it(dialect) mention the driver that from which database we have to connect.

** driver is a raw code written which use all orm including sequelize to connect with particular database.

**Explain models** (makes interoperable between js/ts and mysql)

    Inside the models folder we write the representation(in the form of classes or interface) of the mysql table in javascript/typescript because javascript don't know the sql table, so here models will help us to interect, with mysql database by writing the js like code.

**Explain seeders**(change the data inside the table)

    Inside the seeders folder we put the dummy/seed data of our database, so that whenever a new developer come they can understand the database or tables easily with that dummy data.

**Explain migration**(change the stracture of the database)

    Migration folder is used to create different versions of the database.

    Suppose the airbnb hotel service management system in which initially there were only two columns in a table name and address, now as the product grow and the requirement also grow, we need to add 'STATUS' table as well so instead of update the table we manage the versions of the database so that we can move from V1 TO V2 TO V3 and also can revert back if needed.


**Deleting all the folders created by the sequlieze cli, and organize it in src folder. but why???**

By default it's create config.json file inside all the database credentials present and in the production grade application it's a security compromise that's why we delete all the folder and created config.ts file in config folder where now we can import our db credentials from .env file and rest of the folder like models,seeders and migrations put inside the db folder and make the folder stracture clean. 

## Define .sequelizerc file ##

*put your .sequelizerc file inside the HotleServide* otherwise sequelize cli will not be able to read you .sequelizerc file.

Inside the .sequelizerc file we write the configuration of the sequlize which is automatically picked up by the sequlize cli.

It allows us to define settings like the environment, configuration file path, and paths to migrations, seeders, and models folders.


After creating the files and folders imported the all credentilas from .env file to the config.ts file now we are ready to move further.

## Now we create Migration ##
 
Run command given below, it will create a migration file
```
npx sequelize-cli migration:generate --name create-hote-table
```

**Understand working of Migration**

Migrations has two parts --->> 1.UP, 2.DOWN

1. Up part of the migration contains the code which makes new changes inside our database after running the migration. Below is that code in up part of migration.
```
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
  },
  ```


2. Down part of the migration contains the code which will revert the changes made by the migration if we want to rollback.Below is that code in up part of migration.

```
async down(queryInterface) {
    await queryInterface.sequelize.query(`DROP TABLE IF EXISTS hotels;`);
}

```

here i face some issue while integrating the migration.

By default it support js but we are trying to do it with ts so i got some error
so, for today we execute the migration in js.

**Execute the migration**

To execute the migration run below command

```
npx sequelize-cli db:migrate
```

To revert back the current migration run the below command

```
npx sequelize-cli db:migrate:undo
```

## Benefits of migrations ##

1. Database migration helps use to sequentely update our RDBMS. Because rdbms has strict schema so with the help of migration we can manage our database schmea or versions with time and also we can revert back to the previous versions of database or schema using 'DOWN' part of migration. 

2. If we feel like my previous version of database were more stable and useful according to our business requirement then we can revert back easily to that version.






