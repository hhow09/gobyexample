// Fan in pattern:
// is a concurrency paradigm where inputs from several sources get converged (multiplexed) into a single stream.
// In simple words , this paradigm can be thought of as producer and consumer architecture

// In this example there is 2 ways to merge the channel 1) with sync.WaitGroup 2) with atomic.Add

// Ref <br/>
// 1) Youtube: [Learning Golang: Concurrency Patterns Fan-In and Fan-Out](https://youtu.be/x6vBvgKGvxU) <br/>
// 2) [code](https://github.com/MarioCarrion/videos/tree/3107ff408e0db59b5e9ae07412460375aeb8786a/2021/08/19) <br/>
// 3) [merge channel](https://gist.github.com/Xeoncross/3e0328137019b14373ee26701a23ed81) <br/>
// 4) [Why does my code work correctly when I run wg.Wait() inside a goroutine?](https://stackoverflow.com/questions/46560204/why-does-my-code-work-correctly-when-i-run-wg-wait-inside-a-goroutine)

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

func main() {
	// Read 2 CSV files concurrently. <br/>
	// Output within each file remains in order.\n The result would be produced concurrently.
	p1, _ := filepath.Abs("./assets/file1.csv")
	ch1, err := read(p1)
	if err != nil {
		panic(fmt.Errorf("could not read file %v", err))
	}
	p2, _ := filepath.Abs("./assets/file2.csv")
	ch2, err := read(p2)
	if err != nil {
		panic(fmt.Errorf("could not read file %v", err))
	}
	exit := make(chan struct{})
	// merge 2 channels
	chM := mergeWait(ch1, ch2)
	println("[LOG] receive the merged channel.")
	// chM := mergeAtomic(ch1, ch2)

	go func() {
		for l := range chM {
			fmt.Printf("[CSV Result] %s \n", l)
		}
		close(exit)
	}()

	println("[LOG] Wait for the exit signal.")
	<-exit
	fmt.Println("All completed, Exit.")
}

// Opt 1) merge channels with sync.WaitGroup
func mergeWait(cs ...<-chan []string) <-chan []string {
	out := make(chan []string)
	var wg sync.WaitGroup

	// 1) wait for all channels
	wg.Add(len(cs))

	// 2) mark done after receiving all values from source chan
	send := func(ch <-chan []string) {
		defer wg.Done()
		for v := range ch {
			out <- v
		}
	}

	for _, ch := range cs {
		go send(ch)
	}

	// XXX: fatal error: all goroutines are asleep - deadlock!
	// we need to return out channel first otherwise this function will block the main thread [4]
	//// wg.Wait()
	//// close(out)
	//
	// 3) wait and close channel
	go func() {
		wg.Wait()
		// close channel to avoid deadlock
		close(out)
		println("[LOG] merged channel closed.")
	}()

	return out // 1. return out first
}

// option 2) merge channels with atomic counter
func mergeAtomic(cs ...<-chan []string) <-chan []string {
	out := make(chan []string)
	var counter int32
	atomic.StoreInt32(&counter, int32(len(cs))) // works like wait group
	send := func(c <-chan []string) {
		for v := range c {
			out <- v
		}
		if atomic.AddInt32(&counter, -1) == 0 {
			close(out)
		}
	}
	for _, c := range cs {
		go send(c)
	}
	return out
}

// utils: read file and return a channel of file lines
func read(relPath string) (chan []string, error) {
	absPath, _ := filepath.Abs(relPath)
	f, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("opening file %v", err)
	}
	ch := make(chan []string)
	cr := csv.NewReader(f)

	go func() {
		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(ch)
				return
			}
			ch <- record
		}
	}()
	return ch, nil
}
