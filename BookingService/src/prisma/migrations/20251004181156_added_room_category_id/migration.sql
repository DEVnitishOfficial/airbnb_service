/*
  Warnings:

  - Added the required column `roomCategoryId` to the `Booking` table without a default value. This is not possible if the table is not empty.
  - Made the column `checkInDate` on table `booking` required. This step will fail if there are existing NULL values in that column.
  - Made the column `checkOutDate` on table `booking` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE `booking` ADD COLUMN `roomCategoryId` INTEGER NOT NULL,
    MODIFY `status` ENUM('PENDING', 'CONFIRMED', 'CANCELLED', 'COMPLETED', 'CHECKED_IN', 'CHECKED_OUT') NOT NULL DEFAULT 'PENDING',
    MODIFY `checkInDate` DATETIME(3) NOT NULL,
    MODIFY `checkOutDate` DATETIME(3) NOT NULL;
