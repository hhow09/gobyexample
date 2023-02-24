// ```
// requirement
// A web service provides weather API: GET http://example.server/weather?date=2021-06-29
// query parameter: date: string of date
// rsponse data is of JSON format:
//	{
//	  "status": 0,
//	  "result": {
//		"high": 30,
//		"low":  20,
//		"rain": 0.75
//	  }
//	}
//
// Unfortunately, it does not provide APIs to get a range of data.
// Please complete func quoteApi() so it can take year and month as arguments,
// returning a slice of results.
//
// Bonus: sending requests concurrency
// ```

package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var (
	month int
	year  int
)

type Weather struct {
	High float64 `json:"high"`
	Low  float64 `json:"low"`
	Rain float64 `json:"rain"`
}

const API = "http://example.server/weather"
const TIME_LAYOUT = "2006-1"

// the last day for day count of month
func getLastDay(year int, month int) (int, error) {
	t, err := time.Parse(TIME_LAYOUT, fmt.Sprintf("%d-%v", year, month))
	lastDay := t.AddDate(0, 1, 0).Add(-time.Nanosecond)
	if err != nil {
		return -1, err
	}
	return lastDay.Day(), nil
}

// get single day weather struct from API
func getWeather(ctx context.Context, api string, year, month, day int) (Weather, error) {
	// %02d -> ensure the total lenght of 2 with 0 padding prefix
	url := fmt.Sprintf("%s?date=%d-%02d-%02d", api, year, month, day)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Weather{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Weather{}, err
	}
	if res.StatusCode != 200 {
		return Weather{}, fmt.Errorf("Non-OK HTTP status: %v", res.StatusCode)
	}
	defer res.Body.Close()
	var w Weather
	err = json.NewDecoder(res.Body).Decode(&w)
	if err != nil {
		return Weather{}, err
	}
	return w, nil
}

func quoteApi(api string, year int, month time.Month) ([]Weather, error) {
	lastDay, err := getLastDay(year, int(month))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ret := make([]Weather, lastDay)
	var wg sync.WaitGroup
	var start int32 = 0
	// the success req count
	counter := atomic.LoadInt32(&start)
	for day := 1; day <= lastDay; day++ {
		wg.Add(1)
		go func(ctx context.Context, day int) {
			w, err := getWeather(ctx, api, year, int(month), day+1)
			if err != nil {
				fmt.Println(err)
				// cancel other waiting request if meet one error
				cancel()
				wg.Done()
				return
			}
			atomic.AddInt32(&counter, 1)
			ret[day-1] = w // assign
			wg.Done()
		}(ctx, day)
	}
	wg.Wait()
	// if only got partial data
	if int(counter) != lastDay {
		cancel()
		return nil, errors.New("incomplete data")
	}
	return ret, nil
}

func init() {
	flag.IntVar(&year, "y", 0, "year value")
	flag.IntVar(&month, "m", 0, "month value")
}

func main() {
	flag.Parse()
	quoteApi(API, year, time.Month(month))
}
