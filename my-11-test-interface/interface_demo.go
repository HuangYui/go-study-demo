package main

import "fmt"

type Animal interface {
	eat()
	say()
}

type Dog struct{}

func (Dog) eat() {
	fmt.Printf("dog eat shit")
}

func (Dog) say() {
	fmt.Println("wang wang wang")
}

type Cat struct{}

func (c Cat) eat() {
	println("cat eat 罐头")
}

func (c Cat) say() {
	println("cat say mio")

}

func main() {
	var animal Animal = Dog{}
	animal.say()

	animal = Cat{}
	animal.say()
}
