$ go run http-roundtripper.go
[Sun Jan 22 16:31:10 2023] GET http://httpbin.org/basic-auth/bob/pwd

--- RESPONSE ---
STATUS CODE:  200
BODY:  {
  "authenticated": true, 
  "user": "bob"
}

[Sun Jan 22 16:31:11 2023] GET http://httpbin.org/status/500
[Sun Jan 22 16:31:12 2023] GET http://httpbin.org/status/500
[Sun Jan 22 16:31:14 2023] GET http://httpbin.org/status/500

--- RESPONSE ---
STATUS CODE:  500
BODY:
