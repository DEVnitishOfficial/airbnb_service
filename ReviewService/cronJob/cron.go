package cronjob

import (
	"ReviewService/services"
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func StartCron(svc services.ReviewBatchProcessor, mode string) {
	// starts a new cron job with UTC timezone regardless of server timezone
	c := cron.New(cron.WithLocation(time.UTC))

	var schedule string
	if mode == "prod" {
		schedule = "@hourly"
	} else {
		schedule = "@every 30s"
	}

	_, err := c.AddFunc(schedule, func() {
		// create a context which will expire in 1 minute it avoids long running cron jobs
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		// ctx is a contex object, if we pass this context to any function then they must have to complete in a specified time.
		// using cancel method we can cancel the context before the timeout expires.
		defer cancel() // when function scope exit then it executes

		// we pass ctx to ProcessPendingRatings i.e it must have to implement before one minute, otherwise any goroutine running this function will be aboreted
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
