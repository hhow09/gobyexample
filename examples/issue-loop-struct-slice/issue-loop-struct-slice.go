// Challenge:
//
// What is expected output?
//
// How to fix this code?
//
// Reference: [Go Wiki: Range Clauses](https://go.dev/wiki/Range)
//
// Reference: [Should you use slices of pointers to structs?](https://www.willem.dev/articles/slice-of-pointers-structs/)
//
// Reference: [Go Spec: For statements](https://go.dev/ref/spec#For_statements)

package main

import "fmt"

func main() {
	type player struct {
		name  string
		score int
	}
	players := []player{
		{"Bob", 10},
		{"Alice", 207},
	}

	for _, p := range players {
		p.score += 10
	}

	for _, p := range players {
		fmt.Printf("player %s has score %d\n", p.name, p.score)
	}

	// (End of Challenge)
	// ----------

	// Explaination:
	//
	// - The player struct is stored in the backing array of the slice players
	//
	// - p is declaredin in range clause using a form of [short variable declaration (:=)](https://go.dev/ref/spec#Short_variable_declarations). The variables have the types of their respective iteration values.
	//
	// - `p` is only a copy of the original player
	//
	// - when updating `p.score`, it only updates the "copy" of the player in the slice
	for _, p := range players {
		p.score += 10
	}

	// How to fix it?
	fmt.Println("---- After fix ------")
	// Directly access the player in the slice
	for i := range players {
		players[i].score += 10
	}

	for _, p := range players {
		fmt.Printf("player %s has score %d\n", p.name, p.score)
	}
}
