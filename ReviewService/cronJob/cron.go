package cronjob

import (
	"ReviewService/services"
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func StartCron(svc services.ReviewBatchProcessor, mode string) {
	c := cron.New(cron.WithLocation(time.UTC))

	var schedule string
	if mode == "prod" {
		schedule = "@hourly"
	} else {
		schedule = "@every 30s"
	}

	_, err := c.AddFunc(schedule, func() {
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
	log.Printf("⏰ Cron started with schedule: %s\n", schedule)

	// Run once at startup
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Initial run of ProcessPendingRatings...")
		svc.ProcessPendingRatings(context.Background())
	}()
}
