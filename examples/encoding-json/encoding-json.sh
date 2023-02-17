$ go run encoding-json.go

true marshalled:  true
int 1 marshalled:  1
float 2.340000 marshalled: 2.34
gopher marshelled [34 103 111 112 104 101 114 34]
slcD ["apple","peach","pear"]
mapD {"apple":5,"lettuce":7}
res1D {"Page":1,"Fruits":["apple","peach","pear"]}
res2B {"page":1,"fruits":["apple","peach","pear"]}
=== json.Unmarshal === 
raw json: {"num":6.13,"strs":["a","b"]}, unmarsheled to: map[num:6.13 strs:[a b]]
6.13
a
unmarshel raw json: {"page": 1, "fruits": ["apple", "peach"], "will": "ignore"} to struct: 
{1 [apple peach]}apple
unmarshel raw json: 
{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"} to struct: 
{battery sensor 40 2019-01-21 19:07:28 +0000 UTC}
time field in milli 1548097648000
4. encode to writer: 
{"apple":5,"lettuce":7}