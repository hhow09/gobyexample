# nil

## Overview
- A `nil` slice value needs no allocation. 
- `nil` is a predeclared identifier representing the zero value for a `pointer`, `channel`, `func`, `interface`, `map` or `slice` type
- `nil` is not a keyword

## Kinds of `nil`
- pointers: point to nothing
- slices: have no backing array
- maps: not initialized
- channels: not initialized
- functions: not initialized
- interfaces: have no value assigned, not even a nil pointer


## Nil vs Empty Slice

### Nil slice
- nil and empty slices (with 0 capacity) are not the same, but their observable behavior is the same (almost all the time).
```
[]byte  nil slice
┌──────┬───────┬───────┐
│      │       │       │
│ nil  │   0   │   0   │
└──────┴───────┴───────┘
  ptr     len     cap
```

### Empty Slice
- has underline array of 0 length
```
[]byte  emptynullslice
┌──────┬───────┬───────┐
│      │       │       │
│<add> │   0   │   0   │
└───┬──┴───────┴───────┘
ptrt│     len     cap
    │
    │
┌───▼──┐
│      │
└──────┘
array [0]byte
```

### use case: empty JSON array `[]`
- to marhsal to `[]`, you should create the empty slice (`[]string{}` or `make([]string, 0)`)
- nil slice will return `null`

## nil data structures
### Pointer
```go
var p *int
p == nil // true
*p       // panic: invalid memroy address
```

### slices
```go
var s []slice
len(s)          //0
cap(s)          //0
for range s     // iterates 0 times
s[i]            // panic: index out of range
append(s, i)    // OK
```

### maps
- use `nil` for read-only empty map
```go
var m map[int]int
len(m)          //0
for range n     // iterates 0 times
v, ok := m[i]   // zero(u), false
m[i] = x        // panic: assignment to entry in nil map

func doSomething(options map[string]string) {
    ...
}
// use nil
doSomething(nil)
// instead of 
// doSomething(map[string]string{})
```

### Func
- use `nil` to imply default value in function

### channels
see [channel: merge](../examples/channel-merge/channel-merge.go)

## Ref
- [GopherCon 2016: Francesc Campoy - Understanding nil](https://www.youtube.com/watch?v=ynoY2xz-F8s&ab_channel=GopherAcademy)