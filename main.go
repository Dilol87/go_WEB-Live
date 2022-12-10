package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Map [30][75]int

func main() {

	rand.Seed(time.Now().UnixNano())

	for y := range Map {
		for x := range Map[y] {
			Map[y][x] = rand.Intn(3)

		}
	}

	for y := range Map {
		for x := range Map[y] {
			fmt.Print(Map[y][x])
			if x == (len(Map[y]) - 1) {
				fmt.Println()

			}
		}
	}
}
