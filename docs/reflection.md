# Reflection
## Overview
- interface values is composed of (value, type)
- Reflection is a type of meta programming 
- any interface values can be boxed into reflect.Value and reflect.Type
## Laws of Reflection
### 1. Reflection goes from interface value to reflection object.
### 2. Reflection goes from reflection object to interface value.
### 3. To modify a reflection object, the value must be settable.
```
func (v Value) CanSet() bool
    CanSet reports whether the value of v can be changed. A Value can be
    changed only if it is addressable and was not obtained by the use of
    unexported struct fields. If CanSet returns false, calling Value.Set or any
    type-specific setter (e.g., Value.SetBool, Value.SetInt) will panic.
```
- addressable: need to be pointer variable


## Kind

Although there are infinitely many types, there are only a finite number of kinds of type: the basic types Bool, String and all the numbers; the aggregate types `Array` and `Struct`, the reference types `Chan`, `Func`, `Ptr`, `Slice` and `Map`; `interface` types; and finally Invalid, meaning no value at all (The zero value of a reflect.Value has kind Invalid).

## addressable value
- the address operation `&x` generates a pointer of type `*T` to x. The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation
### not addressable
```go
&m["key"]     // values in a map
&afunc()        // return values from function
&t.method()  // method calls
```


### addressable
```go
v := afunc()
&v
```

## Reference 
- [The Go Blog: The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
- [Go (Golang) Reflection Tutorial](https://youtu.be/f4aUrm7N5DU)
