package main

import "fmt"

func main() {
	i := deferReturn()
	fmt.Println(i)
}

func deferReturnCall() int {
	fmt.Println("deferReturnCall execute")
	return 0
}

func deferReturnCall2() int {
	fmt.Println("return  execute")
	return 6
}

func deferReturn() int {
	defer deferReturnCall()
	return deferReturnCall2()
}
