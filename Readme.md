

# **Airbnb Microservices – Core Features Summary**

This document summarizes the core features and architecture of our Airbnb-style microservices platform.

---

## **1. API Gateway (Golang)**

**Purpose:** The API Gateway acts as the **central entry point** for all requests to the microservices. It ensures **security, routing, and request management**.

### **Key Features:**

* **Central Authentication & Authorization:**

  * Every request to a microservice passes through the gateway.
  * Users must be authenticated via **JWT tokens**.
  * **Role-Based Access Control (RBAC)** ensures only users with appropriate roles (`user` or `admin`) can access certain routes.
  * This protects internal microservices from unauthorized access.

* **Reverse Proxy:**

  * The gateway forwards requests to the correct microservice, hiding the internal endpoints from clients.
  * For example, requests to the BookingService, HotelService, or ReviewService all pass through the gateway first.

* **Rate Limiting:**

  * Requests are limited to **5 per minute per IP** to prevent abuse and ensure fair usage.

* **Internal Microservice Communication:**

  * Two way to communicate with different microservice 

    -  I. Synchronous Communication (Request/Response) 
        * Here we have two options

            i. RESTful APIs -->> use http/https protocol(**used in airbnb**)

            ii. gRPC (Remote Procedure Call) --->> use HTTP/2 with Protocol Buffers (Protobuf) serialization.
    
    - II. Asynchronous Communication (Message-Based)
        * Here also we have two options 

            i. Message Queue method ---->> use AMQP, or Advanced Message Queuing Protocol(RabbitMQ, AWS SQS)**used io-redis and bullmq**

            ii. Publish/Subscribe (Event-Driven) method ----> example Apache Kafka, Redis Pub/Sub

  * The gateway can also handle internal API calls between microservices.
  * It authenticates these calls and aggregates data when needed (e.g., combining user info from AuthService with ReviewService).

**In short:** The gateway acts as a **security checkpoint, traffic manager, and data aggregator** for all microservices.

---

## **2. Booking Service**

**Purpose:** Manages hotel bookings and ensures consistency across the system.

### **Core Logic:**

* Receives booking requests and validates **room availability** via the HotelService(*sync rest api call to hotel microservice*).

* Uses **Redis + RedLock** for distributed locking to prevent multiple users from booking the same room simultaneously.--->> (*await redlock.acquire([bookingResource], ttl);*) ---->bookingREsource has roomId on which lock is put for 60sec after that lock is released.

* Bookings remain in a **pending state** until confirmed.

* After booking creation bookingId is set into the rooms due to which no other user can book that room, but it valid only for 10 mins if the user who created the booking not confirm then when time expires a cronJob runs every minutes who check is there any booking whose status is pending and expiredAt property is less than the current time, then update the booking status as expired and call the hotelService(route /release) to remove the bookingId from the room(make rooms available for other).

* Implementd corner case like the user who has created booking is the same user who is confirming the bookings.


### **Booking Confirmation:**

* Uses **database transactions** to maintain data consistency:

  1. Checks if the operation was already processed using an **IdempotencyKey**.
  2. Updates the booking status to `CONFIRM`.
  3. Marks the IdempotencyKey as finalized to prevent double bookings.

**Key Tables:**

* `Booking` – stores booking info (user, hotel, dates, guests, status).
* `IdempotencyKey` – ensures operations like booking or payment are **idempotent**.

---

## **3. Hotel Service**

**Purpose:** Handles hotel and room management.

### **Key Features:**

1. **Hotel CRUD Operations:** Create, read, update, and delete hotel records.
2. **Room Management:**

   * Check room availability for a given date range.
   * Update room booking IDs when a room is booked.
   * Release rooms for expired or unconfirmed bookings (avoiding “ghost bookings”).
3. **Room Generation & Scheduling:**

   * Bulk creation of rooms using **Redis queues and workers**.
   * Scheduled tasks (cron jobs) extend room availability to maintain a fixed booking window (e.g., 90 days).

---

## **4. Notification Service**

**Purpose:** Sends email notifications to users.

### **Key Features:**

* **Email Queueing:**

  * When an email needs to be sent, it is added to a **Redis queue** with details like recipient, template, and parameters.
* **Workers:** Background workers process the queue and send emails using **nodemailer** with **Handlebars templates** for formatting.

---

## **5. Review Service (Async Rating Calculation)**

**Purpose:** Manages user reviews and calculates hotel ratings.

### **Workflow:**

* A **cron job** periodically aggregates new reviews to calculate hotel ratings.
* Combines review data with **user info from AuthService** to present full review details.
* Updates the HotelService with the **new average rating**.

**Note:** Internal API calls are authenticated through the API Gateway using JWT tokens.

---

## **6. Cross-Cutting Concepts**

* **Idempotency:** Ensures operations like bookings and payments are **not duplicated**.
* **Transactions:** Database operations are executed together to maintain consistency.
* **Redis + RedLock:** Used for **distributed locking** to avoid conflicts in concurrent operations.
* **Cron Jobs:** Scheduled tasks for room availability, cleanup of expired bookings, and asynchronous rating calculations.
* **DTOs & Repositories:** Clean separation of logic layers for **maintainable code**.

---

## **7. Tech Stack & Design Choices**

* **API Gateway:** Golang (fast, concurrent, strongly typed).
* **Microservices:** Node.js + TypeScript for business logic.
* **Database:** Sequelize, Prisma + SQL.
* **Async Jobs:** Redis + BullMQ for email sending and background processing.
* **Security:** JWT + RBAC for authentication and authorization.

---

✅ **Summary:**
This architecture ensures that **all microservices are secured, scalable, and maintainable**. The **API Gateway** centralizes authentication, routing, and request management, while each microservice focuses on its domain logic.

---
