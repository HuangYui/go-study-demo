package main

import "fmt"

const (
	name float32 = iota
)

const (
	a, b = iota, iota
	c, d
	e, f, g = iota, iota, iota
	h, i, j
)

func main() {
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d)
	fmt.Println()
}
