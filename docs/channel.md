# Channel

## Properties of channels
1. Used for synchronization and communication between go routines without explicit locks or condition variables
Internally, works like a FIFO circular queue
2. Channels transfer the copy of the object.
3. By default, sends and receives block until the other side is ready. This allows go-routines to synchronise without explicit locks or condition variables.
4. Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
5. zero-value of a channel is nil
6. When a go-routine G1 wants to receive data from another go-routine G2, but G2 never sends the data, then the channel will make G1 to wait indefinitely, and vice versa.
7. If the buffer is full or if there is nothing to receive, a buffered channel will behave very much like an unbuffered channel.
8. For unbuffered channel, one go-routine should be in running state, while other go-routine should be in runnable state.

## Memory safety
1. the send and receive action need to acquire the lock on the channel
2. The only shared memory between go-routines access is `hchan` which is protected by `mutex`.
3. the task is **copied** from/to the channel 

## hchan struct
```go
type hchan struct {
    qcount   uint      // total data in the queue
    dataqsiz uint      // size of the circular queue
    buf      unsafe.Pointer // pointer to an array(queue), and is nil for unbuffered channel
    elemsize uint16
    closed   uint32    // if channel is closed
    elemtype *_type    // element type
    sendx    uint      // send index
    recvx    uint      // receive index
    recvq    waitq     // list of recv waiters (doubly linked list)
    sendq    waitq     // list of send waiters (doubly linked list)
    lock     mutex     // mutex for concurrent access to the channel
}
```


## Ref
- [Diving Deep Into The Golang Channels.](https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8)
- [Internals of Go Channels](https://shubhagr.medium.com/internals-of-go-channels-cf5eb15858fc)
- [Golang — Understanding channel, buffer, blocking, deadlock and happy groutines.](https://gist.github.com/YumaInaura/8d52e73dac7dc361745bf568c3c4ba37#deadlock)