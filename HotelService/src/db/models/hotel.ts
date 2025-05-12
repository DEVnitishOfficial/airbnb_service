// This file defines a Hotel model using Sequelize ORM for a MySQL database.
// It includes the model's attributes, their types, and some configurations like table name and timestamps.

// Here model is generic class and it can be of different type(like InferAttribute, InferCreationAttribute,UserAttributes, UserCreationAttributes),it's depending on the code we are calling on that time we decide the type of model.

// What is declare keyword ???
// when we are using class then generally we define property inside the constructor but here we are extending the sequelize Model and make our own Hotel class and added extra properties(id,name,address,location,rating...) and we have not use like(this.id=id, this.name=name....)etc. so here using the declare keyword we are saying to the ts-compiler that the given properties(id,name,location,rating....etc) exist and will be provided by the sequelize but i am not assigning it here.

// What is inferAttributes and InferCreationAttributes ???
//InferAttributes<Hotel> automatically infers all attributes (columns) of the Hotel model as they exist in the database.

//InferCreationAttributes<Hotel> infers only the attributes needed when creating a new Hotel instance, respecting CreationOptional.

// ***declare keyword is used to declare variables, properties, or functions without providing an implementation.***

import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";

import sequelize from "./sequelize";

class Hotel extends Model<InferAttributes<Hotel>, InferCreationAttributes<Hotel>> {
    declare id: CreationOptional<number>;
    declare name: string;
    declare address: string;
    declare location: string;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date | null>;
    declare rating?: number;
    declare ratingCount?: number;
}

Hotel.init({
    // Define the attributes of the Hotel model which will be mapped to the database table columns for confirmation see you mysql database table and their columns.
    id: {
        type: 'integer',
        autoIncrement: true,
        primaryKey: true,
    },
    name: {
        type: 'string',
        allowNull: false,
    },
    address: {
        type: 'string',
        allowNull: false,
    },
    location: {
        type: 'string',
        allowNull: false,
    },
    createdAt: {
        type: 'date',
        defaultValue: new Date(),
    },
    updatedAt: {
        type: 'date',
        defaultValue: new Date(),
    },
    deletedAt: {
        type: 'date',
        defaultValue: null,
    },
    rating: {
        type: 'decimal',
        defaultValue: null,
    },
    ratingCount: {
        type: 'number',
        defaultValue: null,
    },
}, {
    tableName: 'hotels', // Specify the table name in the database from which the model will be mapped to. In this case, it is "hotels".
    sequelize: sequelize, // pass the database credentials to the sequelize instance
    underscored: true, // Use snake_case for column names in the database like created_at, updated_at here we have in Hotel model used "createdAt" and "updatedAt" which will be automatically converted into "created_at" and "updated_at" in the database.
    timestamps: true, // Enable timestamps for createdAt and updatedAt fields
});

export default Hotel;