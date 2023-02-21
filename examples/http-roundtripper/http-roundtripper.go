// docs: https://pkg.go.dev/net/http#RoundTripper

// ref: [Go (Golang) http RoundTripper Explained](https://youtu.be/UERCdBrka-o)

// use cases: <br/>
// 1) logging <br/>
// 2) retrying <br/>
// 3) auth <br/>
// 4) caching <br/>
// 5) headers manipulation  <br/>
// 6) testing

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const OK_URL = "http://httpbin.org/status/200"
const BAD_REQ_URL = "http://httpbin.org/status/500"
const NEED_AUTH_URL = "http://httpbin.org/basic-auth/bob/pwd"

func main() {
	// create a chain of roundtripper before
	// the default one (http.DefaultTransport)
	delay := time.Duration(1 * time.Second)
	c := &http.Client{
		Transport: &authRoundTripper{
			next: &retryRoundTripper{
				next: &loggingRoundTripper{
					next:   http.DefaultTransport,
					logger: os.Stdout,
				},
				maxRetries: 3,
				delay:      delay,
			},
			user: "bob",
			pwd:  "pwd",
		},
	}

	// request auth url
	req, err := http.NewRequest(http.MethodGet, NEED_AUTH_URL, nil)
	if err != nil {
		panic(err)
	}
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("\n--- RESPONSE ---")
	fmt.Println("STATUS CODE: ", res.StatusCode)
	fmt.Println("BODY: ", string(body))

	// request status 500 url (should retry)
	req, err = http.NewRequest(http.MethodGet, "http://httpbin.org/status/500", nil)
	if err != nil {
		panic(err)
	}
	res, err = c.Do(req)
	if err != nil {
		panic(err)
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("\n--- RESPONSE ---")
	fmt.Println("STATUS CODE: ", res.StatusCode)
	fmt.Println("BODY: ", string(body))
}

type authRoundTripper struct {
	next      http.RoundTripper
	user, pwd string
}

func (a authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(a.user, a.pwd)
	return a.next.RoundTrip(r)
}

type retryRoundTripper struct {
	next       http.RoundTripper
	maxRetries int
	delay      time.Duration // delay between each retry
}

func (rr retryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var attempts int
	for {
		res, err := rr.next.RoundTrip(r)
		attempts++

		// max retries exceeded
		if attempts == rr.maxRetries {
			return res, err
		}

		// good outcome
		if err == nil && res.StatusCode < http.StatusInternalServerError {
			return res, err
		}

		// delay and retry
		select {
		// return if context is already canceled
		case <-r.Context().Done():
			return res, r.Context().Err()
		case <-time.After(rr.delay):
		}
	}
}

type loggingRoundTripper struct {
	next   http.RoundTripper
	logger io.Writer
}

// RoundTrip is a decorator on top of the default roundtripper
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	// here we can log our message and info
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC),
		r.Method, r.URL.String())
	return l.next.RoundTrip(r)
}
