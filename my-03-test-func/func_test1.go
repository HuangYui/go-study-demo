package main

import "fmt"

func firstFunc(a string, b string) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	c := 100
	return c
}

func secondFunc(a string, b string) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	c := 100
	return c, 200
}

func thirdFunc(a string, b string) (xx int, yy int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	xx = 1000
	yy = 2000
	return
}

func fourthFunc(a string, b string) (xx, yy int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	xx = 1000
	yy = 2000
	return
}

func main() {
	i := firstFunc("Hello", "World!")
	fmt.Println(i)

	i2, i3 := secondFunc("Hello", "World!")
	fmt.Println(i2, i3)

	xx, yy := thirdFunc("asdasd", "asdasd")
	fmt.Println(xx, yy)

}
