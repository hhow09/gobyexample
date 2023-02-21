// Starting with version 1.18, Go has added support for
// _generics_, also known as _type parameters_.

package main

import "fmt"

// `MapKeys` takes a map of any type and returns a slice of its keys.
// This function has two type parameters - `K` and `V`;
// `K` has the `comparable` _constraint_, meaning that
// we can compare values of this type with the `==` and
// `!=` operators. This is required for map keys in Go.
// `V` has the `any` constraint i.e. `interface{}`
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// declare a type constraint as an interface.
type Addable interface {
	int | int32 | int64 | float64
}

// [comparable type](https://go.dev/ref/spec#Comparison_operators)
// comparable: bool, int, float, complex, string, pointer, channel, interface, Array, Struct (if all fields comparable)
// non-comparable: Slice, map, and function types are not comparable.
func SumMapVals[K comparable, V Addable](m map[K]V) V {
	var res V
	for _, v := range m {
		res += v
	}
	return res
}

type LinkList[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// We can define methods on generic types just like we
// do on regular types, but we have to keep the type
// parameters in place. The type is `List[T]`, not `List`.
func (lst *LinkList[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *LinkList[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// When invoking generic functions, we can often rely
	// on _type inference_. Note that we don't have to
	// specify the types for `K` and `V` when
	// calling `MapKeys` - the compiler infers them
	// automatically.
	fmt.Println("keys:", MapKeys(m))

	// ... though we could also specify them explicitly.
	_ = MapKeys[int, string](m)

	intMap := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floatMap := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// generic to sum up the map values
	fmt.Printf("SumMapVals: %v and %v\n",
		SumMapVals(intMap),
		SumMapVals(floatMap))

	lst := LinkList[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("LinkList:", lst.GetAll())
}

// ref: [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics) <br/>
// ref: [Go by Example: Generics](https://gobyexample.com/generics) <br/>
// ref: [comparable type](https://go.dev/ref/spec#Comparison_operators)
