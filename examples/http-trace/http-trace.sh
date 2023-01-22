$ go run http-trace.go
1st Request to example.com

GetConn(example.com:80) 0 ms
DNSStart({Host:example.com}) 0 ms
DNSDone({Addrs:[{IP:2606:2800:220:1:248:1893:25c8:1946 Zone:} {IP:93.184.216.34 Zone:}] 
Err:<nil> Coalesced:false}) 10 ms
ConnectStart(tcp, [2606:2800:220:1:248:1893:25c8:1946]:80)
ConnectDone(tcp, [2606:2800:220:1:248:1893:25c8:1946]:80, <nil>)
GotConn({Conn:0x14000108000 Reused:false WasIdle:false IdleTime:0s}) 153 ms
GotFirstResponseByte 296 ms
PutIdleConn(<nil>)

2nd Request to example.com

GetConn(example.com:80) 0 ms
GotConn({Conn:0x14000108000 Reused:true WasIdle:true IdleTime:72.833µs}) 0 ms
GotFirstResponseByte 142 ms
