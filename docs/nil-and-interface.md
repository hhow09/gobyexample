# nil and interface
## TL;DR
- `nil` represents a zero value in Golang.
- hard-coded `nil` (interface) is always `(type=nil,value=nil)` 

## Interface
An interface contains both the type and the value of the underlying concrete struct that it represents.
```
┌─────────────────────────────────────┐
│             Interface               │
│  ┌───────────────┬───────────────┐  │
│  │     Type      │    Value      │  │
└──┴───────────────┴───────────────┴──┘
```
## pass/assign pointer to interface

### #1 pass to func
```go
func analyzeInterface(i interface{}) {
    fmt.Printf("Interface type: %T\n", i)
    fmt.Printf("Interface value: %v\n", i)
    fmt.Printf("Interface is nil: %t\n", i == nil)
}

var t *A
analyzeInterface(t)
// Interface type: *A
// Interface value: <nil>
// Interface is nil: false
```
because `nil` (`(type=nil,value=nil)` ) has different type with `(type=*A, value=nil)`

### #2 assign nil pointer to interface 
```go
var p *int              // (type=*int,value=nil)
var i interface{}       // (type=nil,value=nil)

if i == nil {           // (type=nil,value=nil) == (type=nil,value=nil)
    fmt.Println("is nil") 
}

i = p                   // assign p to i

// a hardcoded nil is always nil,nil (type,value)
if i != nil {           // (type=*int,value=nil) != (type=nil,value=nil)
    fmt.Println("not a nil")
}

// output:
// is nil
// not a nil
```

## Reference
- [Why Golang Nil Is Not Always Nil? Nil Explained](https://codefibershq.com/blog/golang-why-nil-is-not-always-nil)
- [Understanding nil Interfaces and Interfaces with nil Values in Go](https://trstringer.com/go-nil-interface-and-interface-with-nil-concrete-value/)