# Running this program will cause it to panic, print
# an error message and goroutine traces, and exit with
# a non-zero status.

# When first panic in `main` fires, the program exits
# without reaching the rest of the code. If you'd like
# to see the program try to create a temp file, comment
# the first panic out.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Note that unlike some languages which use exceptions
# for handling of many errors, in Go it is idiomatic
# to use error-indicating return values wherever possible.
