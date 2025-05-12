import { DataTypes, QueryInterface } from "sequelize";

module.exports = {
  // here we can use raw query of mysql and also can use sequelize level query as well.
  async up (queryInterface:QueryInterface) {
    // writing sequelize query, below is parameter of addColumn
    // 1. on which particular table you want to add column
    // 2. name of the column
    // 3. properties of the column
    await queryInterface.addColumn('hotels','deleted_at',{
      type : DataTypes.DATE,
      allowNull : true,
      defaultValue : null,
    })
  },

  async down (queryInterface:QueryInterface) {
    // writing sequelize query, below is parameter of removeColumn
    // 1. on which particular table you want to remove column
    // 2. name of the column
    await queryInterface.removeColumn('hotels','deleted_at')
  }
};

