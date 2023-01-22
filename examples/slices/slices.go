// A _Slices_ is a descriptor of an array segment.
// It consists of a pointer to the array, the length
// of the segment, and its capacity (the maximum length of the segment).

// ref: https://go.dev/blog/slices-intro

package main

import "fmt"

func main() {

	// 1. declare array
	var a [3]int
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3, 4}
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)

	// 2. modify array
	a[2] = 31
	fmt.Println("modify array a", a)

	//
	// 3. slice is just pointer to its underlying array
	// underlying array: [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}
	// underlying array: [6]int{0, 0, 0, 0, 0, 0}
	s2 := make([]int, 3, 6)
	fmt.Println("s1", s1, "cap", cap(s1), "len", len(s1))
	fmt.Println("s2", s2, "cap", cap(s2), "len", len(s2))

	// 4. copy slice with slicing operator
	b1 := s1[1:3]
	fmt.Println("copy slice with slicing operator")
	fmt.Println("b1", b1, "cap", cap(b1), "len", len(b1))

	// all the subsequent slices use the same underlying array under the hood!
	s1[2] = 42
	fmt.Println("b1", b1, "cap", cap(b1), "len", len(b1))

	// 5. copy slice with copy (new memory allocation)
	s1 = []int{1, 2, 3}
	b2 := make([]int, len(s1))
	// same as "copy" func see doc at: go doc builtin.copy
	copy(b2, s1)
	fmt.Printf("copy %v to new array %v\n", s1, b2)
	// or inline copy with append
	fmt.Printf("inline copy %v with append: %v\n",
		s1, append([]int{}, s1...))
	// will not be affected by modifying s1
	fmt.Println("b2", b2, "cap", cap(b2), "len", len(b2))
	s1[2] = 41
	fmt.Println("b2", b2, "cap", cap(b2), "len", len(b2))

	// 6. convert static array to slice
	key := [16]byte{}
	keyS := key[:]
	fmt.Println("keyS", keyS, "cap",
		cap(keyS), "len", len(keyS))

	// 7. increase the length of slice
	s2 = make([]int, 3, 10)
	fmt.Println("s2", s2, "cap", cap(s2), "len", len(s2))
	// increase the length of slice with slicing operator
	// works up until the cap(s2)
	s2 = s2[:5]
	fmt.Println("s2", s2, "cap", cap(s2), "len", len(s2))
	s1 = []int{1, 2, 3}
	fmt.Println("s1", s1, "cap", cap(s1), "len", len(s1))
	s1 = appendInt(s1, 31)
	fmt.Println("s1", s1, "cap", cap(s1), "len", len(s1))

}

func copyInt(dst, src []int) int {
	var n int
	for i, _ := range src {
		dst[i] = src[i]
		n++
	}
	return n
}

func appendInt(s []int, e int) []int {
	expLen := len(s) + 1
	if cap(s) < expLen {
		// grow the slice
		// new allocation create new slice from scratch
		newS := make([]int, expLen, (cap(s)+1)*2)
		// copy the data from old to new
		copyInt(newS, s)
		// assign new to old
		s = newS
	}
	// grow len
	s = s[:expLen]
	// assign the element
	s[expLen-1] = e
	return s
}
