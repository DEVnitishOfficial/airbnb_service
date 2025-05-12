import { QueryInterface } from "sequelize";

// what is QueryInterface
// In Sequelize, QueryInterface is an abstraction provided by Sequelize that allows us to manipulate the database schema directly—without using Sequelize models.

// It is typically used in migrations, where you want to change the structure of your database (e.g., add/remove columns, create/drop tables, etc.).

// In contrast, QueryInterface is for schema manipulation, such as:
// Creating or dropping tables
// Adding or removing columns
// Changing data types
// Creating indexes
// These things affect the structure of your database, not the data itself

// here we can use raw query of mysql and also can use sequelize level query as well

 export default {
  async up (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE hotels
      ADD COLUMN rating DECIMAL(3, 2) DEFAULT NULL,
      ADD COLUMN rating_count INT DEFAULT NULL
      `);
    
  },

  async down (queryInterface:QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE hotels
      DROP COLUMN rating,
      DROP COLUMN rating_count
      `);
  }
};
