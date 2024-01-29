package main

import "fmt"

type Person struct {
	name string
	sex  int
	age  int
}

func (receiver Person) printPerson() {
	println("print start")
	println(receiver.name)
	println(receiver.age)
	println(receiver.sex)
}

func (receiver Person) setName(name string) {
	receiver.name = name
}

func (receiver *Person) setNameRef(name string) {
	receiver.name = name
}
func main() {
	var zhang Person
	zhang.sex = 1
	zhang.name = "张三"
	zhang.age = 66
	fmt.Printf("%v", zhang)

	//第二种方式
	li := Person{
		name: "aaa",
		sex:  20,
		age:  30,
	}

	fmt.Printf("%v", li)
	//方法调用
	zhang.printPerson()
	//值设置（值传递）
	zhang.setName("你是哪个？")
	zhang.printPerson()
	//引用设置
	zhang.setNameRef("你是哪个？")
	zhang.printPerson()

}
