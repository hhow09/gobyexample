// In Go, an _array_ is a numbered sequence of elements of a
// specific length.
// In typical Go code, [slices](slices) are
// much more common; arrays are useful in some special
// scenarios.

package main

import "fmt"

func main() {

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array
	// in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 2 arrays are comparable
	var arr [5]int
	arr[4] = 100
	arr2 := [5]int{0, 0, 0, 0, 100}
	fmt.Println("arr == arr2: ", arr == arr2)

	// copy array
	// re-assign array will be directly copied, and original arr will not be affected
	arr = [5]int{0, 0, 0, 0, 100}
	arr3 := arr
	arr3[4] = 200
	fmt.Println("arr[4] = ", arr[4])

	// reverse array
	arr = [5]int{1, 2, 3, 4, 100}
	rev(&arr) // should pass array with refrence
	fmt.Println("reversed arr = ", arr)
}

func rev(arr *[5]int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}
