package main

import (
	"fmt"
	"math"
)

func main() {

	var a = "initial"

	var b, c = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialfoo

	const s string = "constant"
	const h = 500000000 //5e8
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
