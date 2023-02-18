// Ref: [Why are there nil channels in Go?](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308)
// Ref: [GopherCon 2016: Francesc Campoy - Understanding nil](https://youtu.be/ynoY2xz-F8s)

package main

import "fmt"

// This function creates a channel,  <br/>
// starts a new go routine that sends values to the created channel,  <br/>
// and finally returns the channel.
func buildChan(count int, val int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			c <- val
		}
		close(c)
	}()
	return c
}

// 1st implementation. <br/>
// output: 2 2 1 2 2 2 0 0 0 0 1 0 0 0 0 1 1 0  0 0 0 1 0 0 0 0 0 0 (printing zeros forever)  <br/>
// because we are reading value from closed channel
func merge_wrong1(out chan<- int, a, b <-chan int) {
	go func() {
		for {
			select {
			case v := <-a:
				out <- v
			case v := <-b:
				out <- v
			}
		}
	}()
}

// merge channels 2nd implementation  <br/>
// output: 1 1 1 2 2 1 2 1 2 0 2 0 fatal error: all goroutines are asleep - deadlock!  <br/>
// because we for got to `close(out)`
func merge_wrong2(out chan<- int, a, b <-chan int) {

	go func() {
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					adone = true
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					bdone = true
					continue
				}
				out <- v
			}
		}
	}()
}

// the working case, but bad performance.  <br/>
// infinite loop is reading from case even the channel is closed. <br/>
// it will keep printing "a is now closed" after a is adone
func merge_fixed_1(out chan<- int, a, b <-chan int) {
	go func() {
		defer close(out) // add this
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					fmt.Println("a is now closed")
					adone = true
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					fmt.Println("b is now closed")
					bdone = true
					continue
				}
				out <- v
			}
		}
	}()
}

// the working case.  <br/>
// use nil channel to disable select cases
func merge_fixed_2(out chan<- int, a, b <-chan int) {
	go func() {
		defer close(out)
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					adone = true
					a = nil // add this
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					bdone = true
					b = nil // add this
					continue
				}
				out <- v
			}
		}
	}()
}

func main() {
	ca, cb := buildChan(5, 1), buildChan(5, 2)
	out := make(chan int)
	//// merge_wrong1(out, ca, cb)
	//// merge_wrong2(out, ca, cb)
	//// merge_fixed_1(out, ca, cb)
	merge_fixed_2(out, ca, cb)

	for v := range out {
		fmt.Printf("%v ", v)
	}

}
