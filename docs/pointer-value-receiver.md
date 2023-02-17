# Pointer Receiver v.s. Value Receiver
Should I define methods on values or pointers?

## Overview
- Value method is called with a copy of the caller's argument.
- Pointer receiver passes the address of a type to the function.

## [Should I define methods on values or pointers?](https://go.dev/doc/faq#methods_on_values_or_pointers)
Whether to define the receiver as a value or as a pointer is the same question, then, as whether a function argument should be a value or a pointer. There are several considerations.

1. If the method need to modify the receiver, receiver must be a **pointer**.
2. If the caller is large, and performance is concerned, can use pointer to avoid copy.
3. Value receivers are concurrency safe, while pointer receivers are not concurrency safe. 
4. Try to use the same receiver type for all your methods as much as possible.

## Method Sets
given type T
- the method set of type T consists of all methods with receiver type T
- the method set of type *T consists of all methods with receiver *T or T

## Reference
- https://stackoverflow.com/questions/27775376/value-receiver-vs-pointer-receiver
- https://go.dev/doc/faq#methods_on_values_or_pointers
- https://medium.com/globant/go-method-receiver-pointer-vs-value-ffc5ab7acdb