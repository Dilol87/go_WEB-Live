package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Map [30][75]int

func main() {

	rand.Seed(time.Now().UnixNano())
	correcting_var := (len(Map[0]) - 1)

	for y, m := range Map {
		for x := range m {
			Map[y][x] = rand.Intn(3)

		}
	}

	for y, m := range Map {
		for x := range m {
			fmt.Print(Map[y][x])
			if x == correcting_var {
				fmt.Println()

			}
		}
	}
}
