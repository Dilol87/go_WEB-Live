package main

import (
	conf "GWL/conf"
	funcs "GWL/funcs"
)

func main() {

	var Map [][]int
	Map = funcs.Write_rnd(&Map, conf.RES_Y, conf.RES_X)

	funcs.Draw(Map)
}
