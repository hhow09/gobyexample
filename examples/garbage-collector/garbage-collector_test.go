package main

import "testing"

// func TestMain(t *testing.T) {
// 	for i := 0; i < 100000000; i++ {
// 		d := NewDuck()
// 		birdTalkerInterface(d)
// 	}
// }

func TestMain2(t *testing.T) {
	for i := 0; i < 100000000; i++ {
		d := NewDuck()
		birdTalkerValue(d)
	}
}
