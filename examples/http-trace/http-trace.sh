$ go run http-trace.go
1st Request to http://example.com

GetConn(example.com:80) 0 ms
DNSStart({Host:example.com}) 0 ms
DNSDone({Addrs:[{IP:2606:2800:220:1:248:1893:25c8:1946 Zone:} {IP:93.184.216.34 Zone:}] 
Err:<nil> Coalesced:false}) 11 ms
ConnectStart(tcp, [2606:2800:220:1:248:1893:25c8:1946]:80) 11 ms
ConnectDone(tcp, [2606:2800:220:1:248:1893:25c8:1946]:80, <nil>) 152 ms
GotConn({Conn:0x14000094000 Reused:false WasIdle:false IdleTime:0s}) 152 ms
GotFirstResponseByte (delay) 293 ms
PutIdleConn(<nil>) 294 ms

2nd Request to http://example.com

GetConn(example.com:80) 0 ms
GotConn({Conn:0x14000094000 Reused:true WasIdle:true IdleTime:152.916µs}) 0 ms
GotFirstResponseByte (delay) 140 ms

# for 1st Request, we can see PutIdleConn is called.
# for 2nd Request, DNS resolving part is skipped cuz it reused the previous connection.