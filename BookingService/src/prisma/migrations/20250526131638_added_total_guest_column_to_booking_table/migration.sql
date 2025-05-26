/*
  Warnings:

  - Added the required column `totalGuests` to the `booking` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `booking` ADD COLUMN `totalGuests` INTEGER NOT NULL;