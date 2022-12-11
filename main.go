package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	re_x := 25
	re_y := 10

	var Map [][]int
	Map = write_rnd(&Map, re_y, re_x)

	draw(Map)

}

func draw(Map [][]int) {
	correcting_var := (len(Map[0]) - 1)
	rol := ""

	for y, m := range Map {
		for x := range m {
			rol += fmt.Sprint(Map[y][x])
			if x == correcting_var {
				rol += "\n"
			}
		}
	}
	fmt.Print(rol)
}

func write_rnd(Map_Start *[][]int, y int, x int) [][]int {
	rand.Seed(time.Now().UnixNano())
	Map_S := *Map_Start
	for i := 0; i < y; i++ {
		var buff []int
		for j := 0; j < x; j++ {
			buff = append(buff, rand.Intn(3))
		}
		Map_S = append(Map_S, buff)
	}
	return Map_S
}
