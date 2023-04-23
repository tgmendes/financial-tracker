package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	//ctx := context.Background()
	//
	////limit := rate.NewLimiter(rate.Every(24*time.Hour/time.Duration(50)), 1)
	//burstLimiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(5)), 1)
	//
	//for i := 0; i < 100; i++ {
	//	//if err := limit.Wait(ctx); err != nil {
	//	//	fmt.Println("exceeded max iterations a day")
	//	//
	//	//	return
	//	//}
	//
	//	if err := burstLimiter.Wait(ctx); err != nil {
	//		fmt.Println("exceeded burst rate limit")
	//		return
	//	}
	//
	//	fmt.Printf("[%s] consuming: %d\n", time.Now(), i)
	//}

	gptLimit()
}

func gptLimit() {
	const (
		interval = 60 / 5 * time.Second
		burst    = 5
	)

	limiter := rate.NewLimiter(rate.Every(interval), burst)
	dailyLimiter := rate.NewLimiter(rate.Every(2*time.Minute), 50)
	ctx := context.Background()

	// Allow for up to 50 iterations
	for i := 0; i < 50; i++ {

		// Wait for the limiter to allow the next iteration
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Println("Error waiting for limiter:", err)
			break
		}

		if !dailyLimiter.Allow() {
			fmt.Println("EXCEEDED RATE")
			return
		}

		// Process the iteration
		fmt.Printf("[%s] consuming: %d\n", time.Now(), i)
	}
}
