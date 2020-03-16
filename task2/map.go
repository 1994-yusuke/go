package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	userMap := make(map[string]int)
	rand.Seed(time.Now().UnixNano())
	// TODO
	for i := 0; i <= 10; i++ {
		u, v := game()
		userMap[u] = userMap[u] + v
	}

	var user string
	var score int

	for u, v := range userMap {
		if v > score {
			score = v
			user = u
		}
	}

	fmt.Printf("%v: %v\n", user, score)

}

func game() (user string, score int) {
	switch a := rand.Intn(3); {
	case a == 0:
		return "A", 10
	case a == 1:
		return "B", 10

	default:
		return "C", 10
	}
}
