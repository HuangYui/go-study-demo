package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	fmt.Printf("%T", slice1)
	slice1 = append(slice1, 4)

	for _, b := range slice1 {
		fmt.Printf("%d", b)
	}

	fmt.Printf("%d", cap(slice1))

}
