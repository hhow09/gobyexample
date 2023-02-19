// Types of goroutine leak:
//
// 1) The forgotten sender
//
// 2) The abandoned receiver

// In short:  Never start a goroutine without knowing how it will stop
//
// package to detect goroutine leak [uber-go/goleak](https://github.com/uber-go/goleak)

package main

import (
	"context"
	"fmt"
	"time"
)

// 1) sender with no stopping mechanism
// it does not know when to stop
func forgottenSender() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			fmt.Println("1) blocked ", n)
			ch <- n
			n++
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return ch
}

// 1-fixed) sender with cancel mechanism
func sender_with_cancel(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			// stop sending if canceled
			case <-ctx.Done():
				return
			case ch <- n:
				fmt.Println("1-fixed) blocked ", n)
				n++
				time.Sleep(200 * time.Millisecond)
			}

		}
	}()
	return ch
}

// 2) abandonedWorker
func abandonedWorker(ch chan int) {
	for i := range ch {
		fmt.Println("2) received", i)
	}
}

// worker disaptcher should know when it stop produce <br/>
// but he forgot. <br/>
// fatal error: all goroutines are asleep - deadlock!
func getWork_deadlock() chan int {
	workCh := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			workCh <- i
		}
	}()
	return workCh
}

// 2-fixed) close the channel
func getWork_fixed() chan int {
	workCh := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			workCh <- i
		}
		close(workCh)
	}()
	return workCh
}

func main() {
	fmt.Println("1) forgottenSender is left blocked")
	// 1) forgottenSender is left blocked
	for n := range forgottenSender() {
		fmt.Println("1) received", n)
		if n == 5 {
			break
		}
	}

	fmt.Println("")
	fmt.Println("1-fixed) fixed with context cancel")

	// 1-fixed) fixed with context cancel
	ctx, cancel := context.WithCancel(context.Background())
	for n := range sender_with_cancel(ctx) {
		fmt.Println("1-fixed) received", n)
		if n == 5 {
			cancel() // remember to cancel
			break
		}
	}
	fmt.Println("2) abandonedWorker")
	// 2) abandonedWorker
	// fatal error: all goroutines are asleep - deadlock!
	//// workCh := getWork_deadlock()
	//// abandonedWorker(workCh)

	fmt.Println("2-fixed) abandonedWorker")
	// 2-fixed) abandonedWorker
	workCh_ := getWork_fixed()
	abandonedWorker(workCh_)

}

// ref: https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
//
