package main

import (
	"ReviewService/app"
	clients "ReviewService/client"
	config "ReviewService/config/env"
	"ReviewService/services"
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
)

func main() {
	mode := os.Getenv("APP_MODE") // e.g., "prod" or "test"
	if mode == "" {
		mode = "test" // default to test
	}

	log.Printf("Starting application in %s mode\n", mode)

	// Load config & run app
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)

	// Start HTTP server in a separate goroutine so cron can still run
	go func() {
		app.Run()
	}()

	// Setup DB
	dsn := "root:mysql@1234&?@tcp(localhost:3306)/Airbnb_Review_DB?parseTime=true&loc=UTC"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hotelClient := clients.NewHotelClient("http://localhost:3001/api/v1")
	svc := services.NewReviewBatchProcessor(db, hotelClient)

	// Setup Cron
	c := cron.New(cron.WithLocation(time.UTC))

	var schedule string
	if mode == "prod" {
		schedule = "@hourly"
	} else { // test mode
		schedule = "@every 30s"
	}

	_, err = c.AddFunc(schedule, func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if err := svc.ProcessPendingRatings(ctx); err != nil {
			log.Println("ProcessPendingRatings error:", err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
	log.Printf("Cron started with schedule: %s\n", schedule)

	// Optionally run once at startup
	go func() {
		time.Sleep(2 * time.Second) // small delay to let app start
		log.Println("Initial run of ProcessPendingRatings...")
		svc.ProcessPendingRatings(context.Background())
	}()

	// Block forever (if app.Run() doesn't already do it)
	select {}
}
