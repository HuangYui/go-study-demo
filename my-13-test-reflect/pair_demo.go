package main

import "fmt"

type Man interface {
	eat()
	say()
}

type SuperMan struct {
	name string
	sex  int
}

func (SuperMan) eat() {
	println("man eat")
}

func (SuperMan) say() {
	println("man say")
}

func (rec SuperMan) Get() *SuperMan {
	rec.name = "张三"
	return &rec
}

func main() {
	man := SuperMan{}
	get := man.Get()

	fmt.Printf("%T \n", get)
	fmt.Println(get.name)
	fmt.Println(man.name)
}
