# Dev Notes
## 1. Don’t use Go’s default HTTP client (in production)

no default timeout in default HTTP client

```go
netClient := &http.Client{ Timeout: time.Second * 10 }
response, _ := netClient.Get(url)
```

## 2. Can Check timeout with `err.Timeout()`
```go
if err, ok := err.(net.Error); ok && err.Timeout() {
    ……
}
```

## Ref
- [Don’t use Go’s default HTTP client (in production)](https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779)