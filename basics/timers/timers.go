package main

import (
	"fmt"
	"time"
)

// We often want to execute Go code at some point in the future, or repeatedly at some interval.
// Go’s built-in timer and ticker features make both of these tasks easy.

func main() {

	// Timers represent a single event in the future. You tell the timer how long you want to wait,
	// and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until
	// it sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the
	// timer before it fires. Here’s an example of that.
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// This goroutine, that is waiting on <-timer2.C, remains blocked indefinitely, since timer will
	// be stopped, waiting for a signal on the channel that will never come.

	// This creates a goroutine leak, as the goroutine remains in memory, blocked and unable to
	// proceed or terminate. In real-world applications, such leaks can gradually consume system
	// resources, leading to performance degradation or system instability. To manage such scenarios,
	// it’s usually good to provide a way to unblock such goroutines either through additional
	// synchronization mechanisms (like using a context.Context to signal cancellation) or by
	// ensuring that every channel has a corresponding sender or closer to prevent indefinite waiting.

	// See timers-context.go for an example of how to use a context to cancel a timer.

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.

	time.Sleep(2 * time.Second)

	// The first timer will fire ~2s after we start the program,
	// but the second should be stopped before it has a chance to fire.
}
