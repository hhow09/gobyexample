// 3 ways of singleton pattern

package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	a int
}

var instance *singleton
var mu sync.Mutex

// 1) singleton pattern <br/>
// not thread safe
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// 2) thread safe with lock  <br/>
// check -> lock -> check to reduce the lock access
func GetInstanceThreadSafe() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(singleton)
		}
	}
	return instance
}

// Once is an object that will perform exactly one action.
var once sync.Once

// 3) singleton with once.Do
func GetInstanceThreadSafe2() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func main() {
	i1 := GetInstance()
	i2 := GetInstance()
	fmt.Printf("%p, %p, equal ? %v\n", i1, i2, i1 == i2)

	fmt.Println("GetInstanceThreadSafe")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			p := GetInstanceThreadSafe()
			fmt.Printf("%p ", p)
			wg.Done()
		}()
	}
	wg.Wait()

	instance = nil
	fmt.Println("")
	fmt.Println("GetInstanceThreadSafe2")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			p := GetInstanceThreadSafe2()
			fmt.Printf("%p ", p)
			wg.Done()
		}()
	}
	wg.Wait()
}

// ref: https://blog.kennycoder.io/2021/08/22/Golang-Singleton-%E5%AF%A6%E7%8F%BE%E6%96%B9%E5%BC%8F%E6%8E%A2%E8%A8%8E/
