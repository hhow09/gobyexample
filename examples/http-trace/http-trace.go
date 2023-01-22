// ref: [Go (Golang) httptrace Tutorial](https://youtu.be/u3YWN4TF81w)
//
// ref: [Golang - Request test using net/http/httptrace](https://yushuanhsieh.github.io/post/2019-03-14-http-trace/)

// doc: [Go doc: httptrace](https://pkg.go.dev/net/http/httptrace)

// 1) HTTP Request/Response - Trace Events
//
//  	- DNS Request - DNSStart
//
//  	- DNS Response - DNSDone
//
//  	- TCP Create Connection - ConnectStart/ConnectDone
//
//  	- Write to the TCP connection - WroteHeaders/WroteRequest
//
//  	- Read from TCP Connection - GotFirstResponseByte
//
//  	- Close TCP Connection
//
//
// 2) HTTP Persitent Connections
//
// 		- If TCP Connection already exists use it
//
//  	- Write to the TCP connection - WroteHeaders/WroteRequest
//
//  	- Read from TCP Connection - GotFirstResponseByte
//
//  	- Keep Connection alive for later use - PutIdleConn

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

// tracer for time tracking
type Tracer struct {
	start time.Time
	end   time.Time
}

func (t *Tracer) GetConn(hostPort string) {
	fmt.Printf("GetConn(%s) %d ms\n", hostPort, time.Since(t.start).Milliseconds())
}
func (t *Tracer) DNSStart(info httptrace.DNSStartInfo) {
	fmt.Printf("DNSStart(%+v) %d ms\n", info, time.Since(t.start).Milliseconds())
}
func (t *Tracer) DNSDone(info httptrace.DNSDoneInfo) {
	fmt.Printf("DNSDone(%+v) %d ms\n", info, time.Since(t.start).Milliseconds())
}
func (t *Tracer) ConnectStart(network, addr string) {
	fmt.Printf("ConnectStart(%s, %s) %d ms\n", network, addr, time.Since(t.start).Milliseconds())
}
func (t *Tracer) ConnectDone(network, addr string, err error) {
	fmt.Printf("ConnectDone(%s, %s, %v) %d ms\n", network, addr, err,
		time.Since(t.start).Milliseconds())
}
func (t *Tracer) GotConn(info httptrace.GotConnInfo) {
	fmt.Printf("GotConn(%+v) %d ms\n", info, time.Since(t.start).Milliseconds())
}
func (t *Tracer) GotFirstResponseByte() {
	t.end = time.Now()
	fmt.Printf("GotFirstResponseByte (delay) %d ms\n", time.Since(t.start).Milliseconds())
}

func (t *Tracer) PutIdleConn(err error) {
	t.end = time.Now()
	fmt.Printf("PutIdleConn(%+v) %d ms\n", err, time.Since(t.start).Milliseconds())
}

// Create Trace Info
func createTrace() *httptrace.ClientTrace {
	tracer := Tracer{start: time.Now()}
	trace := &httptrace.ClientTrace{
		GetConn:              tracer.GetConn,
		DNSStart:             tracer.DNSStart,
		DNSDone:              tracer.DNSDone,
		ConnectStart:         tracer.ConnectStart,
		ConnectDone:          tracer.ConnectDone,
		GotConn:              tracer.GotConn,
		GotFirstResponseByte: tracer.GotFirstResponseByte,
		// called when connection is put back to connection pool
		PutIdleConn: tracer.PutIdleConn,
	}
	return trace
}

const url = "http://example.com"

func main() {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	trace := createTrace()

	// Create Trace context
	ctx := httptrace.WithClientTrace(req.Context(), trace)

	// Attach Trace context to request
	req = req.WithContext(ctx)

	fmt.Println("1st Request to", url)
	fmt.Println("")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// read the whole body and close so that the underlying TCP conn is re-used
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	fmt.Println("")
	fmt.Println("2nd Request to", url)
	fmt.Println("")
	req2, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	trace = createTrace()
	// Create Trace context
	ctx = httptrace.WithClientTrace(req2.Context(), trace)

	// Attach Trace context to request
	req = req.WithContext(ctx)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
