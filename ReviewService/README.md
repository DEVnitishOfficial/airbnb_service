# Review Service

A Go-based microservice for managing hotel reviews in the Airbnb system.

## Features

- CRUD operations for reviews
- Filter reviews by user, hotel, or booking
- Soft delete functionality
- Input validation
- RESTful API endpoints

## Database Schema

The service uses MySQL database `airbnb_reviews` with the following table structure:

```sql
CREATE TABLE reviews (
 id BIGINT AUTO_INCREMENT PRIMARY KEY,
 user_id BIGINT NOT NULL,
 booking_id BIGINT NOT NULL,
 hotel_id BIGINT NOT NULL,
 comment TEXT NOT NULL,
 rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 deleted_at TIMESTAMP NULL,
 is_synced BOOLEAN NOT NULL DEFAULT FALSE
);
```

## API Endpoints

### CRUD Operations
- `POST /reviews` - Create a new review
- `GET /reviews` - Get all reviews
- `GET /reviews/{id}` - Get review by ID
- `PUT /reviews/{id}` - Update a review
- `DELETE /reviews/{id}` - Delete a review (soft delete)

### Filter Operations
- `GET /reviews/user?user_id={id}` - Get reviews by user ID
- `GET /reviews/hotel?hotel_id={id}` - Get reviews by hotel ID
- `GET /reviews/booking?booking_id={id}` - Get reviews by booking ID

## Setup

1. **Install dependencies:**
   ```bash
   make deps
   ```

2. **Set up environment variables:**
   Create a `.env` file with:
   ```
   DB_USER=root
   DB_PASSWORD=root
   DB_NET=tcp
   DB_ADDR=127.0.0.1:3306
   DBName=airbnb_reviews
   PORT=:8081
   ```

3. **Run database migrations:**
   ```bash
   make migrate-up
   ```

4. **Run the service:**
   ```bash
   make run
   ```

## Example Usage

### Create a Review
```bash
curl -X POST http://localhost:8081/reviews \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "booking_id": 123,
    "hotel_id": 456,
    "comment": "Great hotel with excellent service!",
    "rating": 5
  }'
```

### Get All Reviews
```bash
curl http://localhost:8081/reviews
```

### Get Reviews by Hotel
```bash
curl "http://localhost:8081/reviews/hotel?hotel_id=456"
```

## Architecture

The service follows a clean architecture pattern:
- **Controllers**: Handle HTTP requests and responses
- **Services**: Business logic layer
- **Repositories**: Data access layer
- **Models**: Data structures
- **DTOs**: Data transfer objects for API requests/responses
- **Middlewares**: Request validation and processing
- **Router**: URL routing and middleware chaining

## Dependencies

- Go 1.24.1+
- MySQL 8.0+
- Chi router for HTTP routing
- Goose for database migrations
- Validator for input validation 


# 🏨 Review Aggregation Cron Service

This microservice is responsible for **periodically aggregating hotel review ratings** and **updating the HotelService** with the latest average rating and review count.
It runs as a **scheduled background job** using a cron mechanism, ensuring that new reviews are processed regularly and hotel ratings stay in sync across services.

---

## 🚀 Overview

The **Review Aggregation Cron Service** automates the process of:

1. Fetching all new (unsynced) reviews from the `reviews` table.
2. Grouping them by `hotel_id` to calculate:

   * Total rating sum
   * Total review count
3. Fetching the hotel’s existing rating and count from **HotelService** (an external microservice).
4. Calculating the **new average rating** using:

   ```
   newAvg = (oldAvg * oldCount + sumOfNewRatings) / newCount
   ```
5. Updating the **HotelService** with the new `rating` and `rating_count`.
6. Marking those reviews as synced (`is_synced = TRUE`) in the local database.

---

## 🧠 Architecture & Layered Design

The project follows a **clean layered architecture** to promote readability, testability, and scalability.

```
ReviewService/
├── main.go                      → Application entry point
├── config/
│   ├── env/                     → Environment variable loading
│   └── db/db.go                 → Database connection (MySQL)
├── repository/
│   └── review_repository.go     → DB queries for reviews (data layer)
├── services/
│   └── review_batch_service.go  → Business logic (aggregation, transactions)
├── client/
│   └── hotel_client.go          → HTTP client to communicate with HotelService
└── cron/
    └── rating_cron.go           → Cron job scheduler logic
```

### 🧩 Layer Responsibilities

| Layer                | Responsibility                                                                                   |
| -------------------- | ------------------------------------------------------------------------------------------------ |
| **Repository Layer** | Handles raw SQL queries for reading and updating reviews.                                        |
| **Service Layer**    | Contains business logic for rating aggregation, locking, and transactional updates.              |
| **Client Layer**     | Makes HTTP requests to the external `HotelService` microservice for reading/updating hotel data. |
| **Cron Layer**       | Uses `github.com/robfig/cron/v3` to schedule periodic executions of the service.                 |
| **Config Layer**     | Loads environment variables and manages MySQL connections securely using `.env` values.          |

---

## 🧮 How It Works (Step-by-Step)

1. **Cron Job Triggered**

   * In **dev mode**, runs every **30 seconds**.
   * In **production mode**, runs **hourly (@hourly)**.

2. **Lock Acquisition**

   * Uses MySQL’s `GET_LOCK()` to prevent multiple instances from running the batch simultaneously.

3. **Aggregate Pending Reviews**

   * Groups reviews with `is_synced = FALSE` up to the current UTC time cutoff.
   * Calculates sum and count per hotel.

4. **Fetch Hotel Data**

   * Calls `GET /api/v1/hotels/:id` on **HotelService** to get existing rating and count.

5. **Compute New Rating**

   * Uses weighted average formula to merge old and new ratings.

6. **Update Hotel**

   * Calls `PATCH /api/v1/hotels/:id` on **HotelService** to update the rating and count.

7. **Mark Reviews as Synced**

   * Updates the processed reviews as `is_synced = TRUE`.

8. **Release Lock**

   * Ensures MySQL lock is released even if the process fails midway.

---

## 🕒 Cron Configuration

We use the **robfig/cron v3** package for scheduling:

```go
import "github.com/robfig/cron/v3"
```

Example configuration:

```go
c := cron.New(cron.WithLocation(time.UTC))

if mode == "prod" {
    schedule = "@hourly"
} else {
    schedule = "@every 30s"
}

c.AddFunc(schedule, func() {
    svc.ProcessPendingRatings(context.Background())
})
c.Start()
```

---

## 🔐 Environment Configuration

Database and environment values are loaded from `.env` using a helper package.

**Example `.env`:**

```
APP_MODE=dev
DB_USER=root
DB_PASSWORD=your_password
DB_NET=tcp
DB_ADDR=127.0.0.1:3306
DB_NAME=Airbnb_Review_DB
HOTEL_SERVICE_URL=http://localhost:3001/api/v1
```

**Example DB Setup (`config/db/db.go`):**

```go
cfg := mysql.NewConfig()
cfg.User = env.GetString("DB_USER", "root")
cfg.Passwd = env.GetString("DB_PASSWORD", "")
cfg.Net = "tcp"
cfg.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
cfg.DBName = env.GetString("DB_NAME", "Airbnb_Review_DB")
cfg.Params = map[string]string{"parseTime": "true", "loc": "UTC"}
```

---

## 🌍 Environment Switching (Dev ↔ Prod)

The app automatically detects mode via:

```go
mode := os.Getenv("APP_MODE")
```

| Mode     | Schedule         | Description                           |
| -------- | ---------------- | ------------------------------------- |
| **dev**  | every 30 seconds | Quick feedback loop for local testing |
| **prod** | hourly           | Optimized for production workloads    |

Simply change `APP_MODE` in `.env` and restart the service — no code change required.

---

## 🧩 Packages Used

| Package                                 | Purpose                                                      |
| --------------------------------------- | ------------------------------------------------------------ |
| `github.com/robfig/cron/v3`             | Scheduling periodic background jobs                          |
| `github.com/go-sql-driver/mysql`        | MySQL driver for database access                             |
| `database/sql`                          | Standard Go SQL interface                                    |
| `context`                               | Timeout and cancellation handling for DB and HTTP operations |
| `net/http`                              | Communication with external HotelService                     |
| `github.com/joho/godotenv` *(optional)* | Load `.env` configuration                                    |

---

## 🧪 Local Testing (Dev Mode)

1. Run MySQL locally and ensure your `reviews` and `hotels` tables exist.
2. Add `.env` file in project root.
3. Start the service:

   ```bash
   go run main.go
   ```
4. Every 30 seconds, new reviews (`is_synced=FALSE`) will be processed and the updated rating pushed to HotelService.

---

## 🏁 Production Deployment

1. Set:

   ```
   APP_MODE=prod
   ```
2. Deploy the service alongside HotelService.
3. The cron will run hourly, and all logs will show UTC timestamps.

---

## 📜 Summary

✅ Clean separation of **Repository**, **Service**, and **Client** layers
✅ Uses **MySQL locks** for safe concurrency
✅ Uses **robfig/cron/v3** for reliable scheduling
✅ Easy environment-based configuration
✅ Communicates with external **HotelService** via REST APIs
✅ Fully decoupled, testable, and maintainable design

---
