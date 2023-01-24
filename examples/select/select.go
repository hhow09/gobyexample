// select is implemented in src/runtime/select.go
// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.”

// Details
//
// 		1) lock all channel in scase syntax
//
//		2) randomly check if the any channel of the case is ready
//
//			2.1 if case can read, then return (case index, true)
//
//			2.2 if case can write, wrtie data to channelm unlock all channels and return (case index, false)
//
//   		2.3 (with default) if all cases are not ready yet, then unlock all channels and return（default index, false）
//
//		3) (without default) if all cases are not ready yet
//
//			3.1 put current goroutine into wait queue
//
//          3.2 block current goroutine until any channel in case is ready.
//
//      4) if ready, go to step 2

package main

import (
	"fmt"
	"time"
)

func main() {
	// 1) For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 2) when multiple cases are available, it will select a random one
	ch := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
	}()
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	}

	// 3) use case "default" to prevent deadlock
	ch2 := make(chan int, 1)
	select {
	case <-ch2:
		fmt.Println("got from ch2")
	case <-ch2:
		fmt.Println("got from ch2")
	default:
		fmt.Println("default: not receiving any value")
	}

	// 4) prevent goroutine blocked forever with timeout "time.After"
	ch3 := make(chan int, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch3 <- 1
	}()
	select {
	case v := <-ch3:
		fmt.Println("receive ch3 value", v)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout!")
	}
}
