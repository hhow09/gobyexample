// [singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight) provides a duplicate function call suppression mechanism.
// you create a key to identify the request. when there're other requests
// with the same key, it will wait for the answer which is processing for
// another request.

package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	g            singleflight.Group
	ErrCacheMiss = errors.New("cache miss")
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("===default query===")
	wg.Add(10)
	// simulating 10 concurrent request
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := loadDefault("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()

	fmt.Println("")
	fmt.Println("===wrap query with singleflight===")
	wg.Add(10)
	// simulating 10 concurrent request
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := loadWithSingleflight("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}

func loadDefault(key string) (string, error) {
	data, err := loadFromCache(key)
	if err != nil && err == ErrCacheMiss {
		data, err := loadFromDB(key)
		if err != nil {
			log.Println(err)
			return "", err
		}
		setCache(key, data)
	}
	return data, nil
}

// load from cache, if missed, load from db and set cache
func loadWithSingleflight(key string) (string, error) {
	data, err := loadFromCache(key)
	if err != nil && err == ErrCacheMiss {
		v, err, _ := g.Do(key, func() (interface{}, error) {
			data, err := loadFromDB(key)
			if err != nil {
				return nil, err
			}
			setCache(key, data)
			return data, nil
		})
		if err != nil {
			log.Println(err)
			return "", err
		}
		data = v.(string)
	}
	return data, nil
}

// simulate cache miss
func loadFromCache(key string) (string, error) {
	return "", ErrCacheMiss
}

// setCache
func setCache(key, data string) {}

// getDataFromDB
func loadFromDB(key string) (string, error) {
	fmt.Println("query db")
	unix := strconv.Itoa(int(time.Now().UnixNano()))
	return unix, nil
}
