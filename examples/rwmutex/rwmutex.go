// ref: [sync.RWMutex Solving readers-writers problems](https://medium.com/golangspec/sync-rwmutex-ca6c6c3208a0)
// Readers-writers problems occur when shared piece of data needs to be accessed by multiple threads.

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func reader(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	sleep()
	m.RLock() // 1) acquire Read lock
	c <- 1    // 2) add acquired lock count
	sleep()   // 3) do something after acquiring lock
	c <- -1   // 4) reduce acquired lock count
	m.RUnlock()
	wg.Done()
}

func writer(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	sleep()
	m.Lock() // 1) acquire Write/Exclusive lock
	c <- 1
	sleep()
	c <- -1
	m.Unlock()
	wg.Done()
}

const READER_COUNT = 10
const WRITER_COUNT = 3

func main() {
	var m sync.RWMutex
	var rs, ws int
	readerCh := make(chan int)
	writerCh := make(chan int)

	// print out the lock count result
	go func() {
		for {
			select {
			case n := <-readerCh:
				rs += n
			case n := <-writerCh:
				ws += n
			}
			// simulate concurrent lock acquired count
			fmt.Printf("%s%s\n", strings.Repeat("R", rs),
				strings.Repeat("W", ws))
		}
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reader(readerCh, &m, &wg)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go writer(writerCh, &m, &wg)
	}
	wg.Wait()
}
