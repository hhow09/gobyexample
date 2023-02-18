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
	s3 := []int{1, 2, 3}
	b3 := make([]int, len(s3))
	// same as "copy" func see doc at: go doc builtin.copy
	copy(b3, s3)
	fmt.Printf("copy %v to new array %v\n", s3, b3)
	// or inline copy with append
	fmt.Printf("inline copy %v with append: %v\n",
		s3, append([]int{}, s3...))
	// will not be affected by modifying s1
	fmt.Println("b3", b3, "cap", cap(b3), "len", len(b3))
	s3[2] = 41
	fmt.Println("b3", b3, "cap", cap(b3), "len", len(b3))

	// 6. convert static array to slice
	key := [16]byte{}
	keyS := key[:]
	fmt.Println("keyS", keyS, "cap",
		cap(keyS), "len", len(keyS))

	// 7. increase the length of slice
	s4 := make([]int, 3, 10)
	fmt.Println("s4", s4, "cap", cap(s4), "len", len(s4))
	// increase the length of slice with slicing operator
	// works up until the cap(s4)
	s4 = s4[:5] // len=3, cap=5
	fmt.Println("s4", s4, "cap", cap(s4), "len", len(s4))
	s5 := []int{1, 2, 3}
	fmt.Println("s5", s5, "cap", cap(s5), "len", len(s5))
	// trigger the capacity increase since reached the threshold
	s5 = append_(s5, 31) // cap 3 -> 8, len 3 -> 4
	fmt.Println("s5", s5, "cap", cap(s5), "len", len(s5))

}

func copySlice[K any](dst, src []K) int {
	var n int
	for i, _ := range src {
		dst[i] = src[i]
		n++
	}
	return n
}

// simulate the append of go internal
func append_[K any](s []K, e K) []K {
	expLen := len(s) + 1
	if cap(s) < expLen {
		// grow the slice
		// new allocation create new slice from scratch
		newS := make([]K, expLen, (cap(s)+1)*2)
		// copy the data from old to new
		copySlice(newS, s)
		// assign new to old
		s = newS
	}
	// grow len
	s = s[:expLen]
	// assign the element
	s[expLen-1] = e
	return s
}
