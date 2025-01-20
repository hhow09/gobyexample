// interview questions:
//
// What is expected output?
//
// How to fix this code?

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

	// (End of Question)
	// ----------

	// Explaination:
	// The player struct is stored in the backing array of the slice players
	//
	// `p` is actually only a copy of the original player
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

	// Output:
	//
	// player Bob has score 20
	//
	// player Alice has score 217
}

// Reference: [Should you use slices of pointers to structs?](https://www.willem.dev/articles/slice-of-pointers-structs/)
