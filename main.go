package main

import (
	conf "GWL/conf"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var Map [][]int
	Map = write_rnd(&Map, conf.RES_Y, conf.RES_X)

	draw(Map)
}

func draw(Map [][]int) {
	correcting_var := (len(Map[0]) - 1)
	rol := ""

	for y, m := range Map {
		for x := range m {
			switch Map[y][x] {
			case 0:
				rol += " "
			case 1:
				rol += "0"
			}

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
			buff = append(buff, rand.Intn(2))
		}
		Map_S = append(Map_S, buff)
	}
	return Map_S
}
