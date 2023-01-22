

go run rwmutex.go
# New line is printed each time set of goroutines executing critical section changes. 
# It’s visible that RWMutex allows for either 
#   1) at least one reader 
#   or
#   2) exactly one writer.
W
R
RR
RRR
RRRR
RRRRR
RRRR
RRR
RRRR
RRR
RR
R
W
R
RR
RRR
RRRR
RRR
RR
R
W

