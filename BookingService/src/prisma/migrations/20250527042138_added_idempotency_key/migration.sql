/*
  Warnings:

  - A unique constraint covering the columns `[IdempotencyKeyId]` on the table `Booking` will be added. If there are existing duplicate values, this will fail.

*/
-- AlterTable
ALTER TABLE `booking` ADD COLUMN `IdempotencyKeyId` INTEGER NULL;

-- CreateTable
CREATE TABLE `idempotencyKey` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `key` INTEGER NOT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL,

    UNIQUE INDEX `idempotencyKey_key_key`(`key`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateIndex
CREATE UNIQUE INDEX `Booking_IdempotencyKeyId_key` ON `Booking`(`IdempotencyKeyId`);

-- AddForeignKey
ALTER TABLE `Booking` ADD CONSTRAINT `Booking_IdempotencyKeyId_fkey` FOREIGN KEY (`IdempotencyKeyId`) REFERENCES `idempotencyKey`(`id`) ON DELETE SET NULL ON UPDATE CASCADE;
