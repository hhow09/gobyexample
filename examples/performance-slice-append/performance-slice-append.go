// Compare and benchmark 3 kinds of slice append <br/>
// ref: [Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices)
// ref: [Golang slice append vs assign performance](https://stackoverflow.com/questions/38654729/golang-slice-append-vs-assign-performance)

package main

const length = 1000

// 1) `genByAppend` will encountered several times of re-allocation as the slice grows. Process includes:<br/>
//   - re-allocate a new array with doubled size <br/>
//   - copy the existing array data to new one <br/>
//   - change the slice pointer to array <br/>
//   - append <br/>
func genByAppend() []int {
	var s []int
	for i := 0; i < length; i++ {
		s = append(s, i)
	}
	return s
}

// 2) `genByAppendCap` pre-allocated the capacity of `length` <br/>
//   - already allocated enough space. <br/>
//   - should not trigger the re-allocation <br/>
//   - however the local variable s (which is a slice header) has to be updated in each cycle of the loop.
func genByAppendCap() []int {
	s := make([]int, 0, length)
	for i := 0; i < length; i++ {
		s = append(s, i)
	}

	return s
}

// 3) `genByAssign` is the fastest among 3, because of directly assignment.
func genByAssign() []int {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = i
	}
	return s
}
