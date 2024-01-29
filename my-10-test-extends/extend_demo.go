package main

import "fmt"

type Human struct {
	name string
	sex  int
}

type PersonA struct {
	Human
	age int
}

func (receiver Human) printName() {
	fmt.Println(receiver.name)
}

func main() {
	zhangsan := PersonA{
		Human: Human{"张三", 1},
		age:   0,
	}
	zhangsan.name = "111"
	zhangsan.printName()
	fmt.Println(zhangsan)
}
