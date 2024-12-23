# Interview Questions
## Data structure
1. 如果range over slice時對slice作改動，是否會改動slice ?
range over slice will copy the original value.
2. how to compare 2 struct value ?
`reflect.DeepEqual`
3. how to verify that struct actually implement a Interface ?
`var _ Interface = (*Struct_Should_Implement)(nil) `
4. convert `[]byte` to `string` 是否會造成 copy ? 
   - yes
   - string is immutable but `[]byte` is mutable
5. Modifying map while iterating over it ?
   - unpredictable
   - If a map entry is created during iteration, that entry may be produced during the iteration or may be skipped.
   - If a map entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced.

6. How to check if a slice is sorted ?
    - [sort.IsSorted()](https://golang.org/pkg/sort/#IsSorted)
    - [sort.SliceIsSorted()](https://golang.org/pkg/sort/#SliceIsSorted), passing a less()
7. marshal json what is the result ?
```go
type J struct {
	a string 
	b string `json:"B"`
	C string
	D string `json:"DD"`
}
	j := J{
		a: "1",
                 b: "2",
		C: "3",
		D: "4",
	}
	fmt.Printf("j = %+v\n", j) 
```
8. make vs new
    - both for memory allocation
    - new will allocate the zero value of type
    - make can only allocate map, slice, chan
9. can we address the map ?
    - map is not addressable
    - when hash table is reorganized, the address will change.
    ```go
    m := make(map[string]int)
    fmt.Println(&m["qcrao"])
    // cannot take the address of m["qcrao"]
    ```

10. use case of `tag`
    - json marshall unmarshal
    - db: sqlx column mapping
    - form: gin validation for form
    - binding: gin binding for validation

11. `fmt.Printf` with `%v` `%+v` `%#v`
```go
type animal struct {
	name string
	age  int
}

a := animal{name: "cat", age: 2}
fmt.Printf("%v\n", a) // {cat 2}
fmt.Printf("%+v\n", a) // {name:cat age:2}
fmt.Printf("%#v\n", a) // main.animal{name:"cat", age:2}
```

## Functions
1. the execution order of import, const, var, init(), main()
import –> const –> var –>init()–>main()

## Concurrency

1. Channel deadlock的情況
    - read from `nil` channel
    - write to full channel
    - write to closed channel
    - all grouting blocked
3. 讀closed channel 的行為? 
    - 一直讀到 zero value
    - 不會deadlock
    - 解決: 把原chan 設為 nil


### Concurrency control
1. use wait group to wait for multiple goroutines returns
2. use channel to block channel sender or channel receiver
4. use channel if we need specific order of goroutine execution
5. pass context within chain of goroutine to control / cancel .
6. use select to organize multiple channels

### Mutex

## Design Pattern

## Internals
### [Channel](https://github.com/hhow09/gobyexample/blob/master/docs/channel.md)

### Map
1. map is `hmap` struct with pointer `buckets` to buckets array of `2^B` buckets.
2. bucket is `bmap` struct with 8 buckets, each bucket stores an key value pair

## Web
1. Graceful shutdown
    1. start a goroutine for `server.listenAndServe()` in background 
    13. main thread blocked and listen to os termination signal (`syscall.SIGTERM` )
    14. if received os signal, use context with timeout and call server to shutdown.


## Ref
- [Go 语言问题集(Go Questions)](https://www.bookstack.cn/books/qcrao-Go-Questions)
- [[吐血整理]超全golang面试题合集+golang学习指南+golang知识图谱+成长路线](https://learnku.com/articles/56078)
- [Go常见面试题【由浅入深】2022版](https://zhuanlan.zhihu.com/p/471490292)
- [你遇到过哪些高质量的 Go 语言面试题？](https://www.zhihu.com/question/60952598)
- [Golang面试问题汇总](http://wearygods.online/articles/2021/04/19/1618823886966.html)
- [GO語言面試系列:（二）常規性Golang面試題解析](https://www.gushiciku.cn/pl/2Oti/zh-tw)
- [34 Go Interview Questions To Crack In 2020](https://www.fullstack.cafe/blog/go-interview-questions)
- [50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs](http://golang50shad.es/index.html#map_value_field_update)
