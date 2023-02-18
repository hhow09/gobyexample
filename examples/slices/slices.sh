# Note that while slices are different types than arrays,
# they are rendered similarly by `fmt.Println`.
a [0 0 0]
b [1 2 3]
c [1 2 3 4]
modify array a [0 0 31]
s1 [1 2 3] cap 3 len 3
s2 [0 0 0] cap 6 len 3
copy slice with slicing operator
b1 [2 3] cap 2 len 2
b1 [2 42] cap 2 len 2
copy [1 2 3] to new array [1 2 3]
inline copy [1 2 3] with append: [1 2 3]
b3 [1 2 3] cap 3 len 3
b3 [1 2 3] cap 3 len 3
keyS [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] cap 16 len 16
s4 [0 0 0] cap 10 len 3
s4 [0 0 0 0 0] cap 10 len 5
s5 [1 2 3] cap 3 len 3
s5 [1 2 3 31] cap 8 len 4

# Check out this [great blog post](https://go.dev/blog/slices-intro)
# by the Go team for more details on the design and
# implementation of slices in Go.