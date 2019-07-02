package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		randInt(1, 1000000000)
	}
}

func randInt(min int, max int) int {
	fmt.Println(rand.Intn(max - min))
	return min + rand.Intn(max-min)
}
