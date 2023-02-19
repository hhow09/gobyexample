$ go run issue-goroutine-leak.go

1) forgottenSender is left blocked
1) blocked  0
1) received 0
...
1) blocked  5
1) received 5

1-fixed) fixed with context cancel
1-fixed) blocked  0
1-fixed) received 0
1-fixed) blocked  1
1-fixed) received 1

// we can see the leaked goroutine of example-1
1) blocked  6       // the forgottne sender

1-fixed) blocked  2
1-fixed) received 2
...
1-fixed) blocked  5
1-fixed) received 5

2) abandonedWorker
2-fixed) abandonedWorker
2) received 0
2) received 1
2) received 2
2) received 3