package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Timers represent a single event in the future.
	timer1 := time.NewTimer(2 * time.Second)

	// Block on the timer's channel until it sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Timer to demonstrate stopping and using context for goroutine cancellation.
	timer2 := time.NewTimer(time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-timer2.C:
			fmt.Println("goroutine executing: Timer 2 fired")
		case <-ctx.Done():
			fmt.Println("goroutine exiting: Timer 2 is stopped")
			return
		}
	}()

	// Stop the timer and cancel the context if the timer is successfully stopped before it fires.
	if stop2 := timer2.Stop(); stop2 {
		cancel()
		fmt.Println("Timer 2 stopped")
	}

	// Wait some time to see the effect of the timer stop.
	time.Sleep(2 * time.Second)
}
