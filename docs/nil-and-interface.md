# nil and interface
## TL;DR
- `nil` represents a zero value in Golang.
- hard-coded `nil` (interface) is always `(type=nil,value=nil)`
- do not return concreate error types

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

## Check if interface holds a nil value

```go
import (
    "reflect"
    "unsafe"
)


// 1) will panic on input struct value
// reflect: call of reflect.Value.IsNil on struct Value
func isNil1(i interface{}) bool{
    return i == nil || reflect.ValueOf(i).IsNil()
}

func isNil2(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	// only works for pointer types
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

```

## Do not return concrete error type

prevent this bad practice

```go
func readFile() error {
    var err *fs.PathError
    fmt.Println(err == nil) // true
    return err
}

func main() {
    err := readFile()
    print(err == nil) // `false` because it is (type=*fs.PathError, value=nil)
}
```
- `error` type is an interface
- should return `nil` instead of `nil pointer`

## Reference
- [Why Golang Nil Is Not Always Nil? Nil Explained](https://codefibershq.com/blog/golang-why-nil-is-not-always-nil)
- [Understanding nil Interfaces and Interfaces with nil Values in Go](https://trstringer.com/go-nil-interface-and-interface-with-nil-concrete-value/)