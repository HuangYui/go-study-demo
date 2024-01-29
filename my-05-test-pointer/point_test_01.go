package main

import "fmt"

// 引用传递
func changeRefValue(a *int) {
	*a = 10
}

// 值传递
func changeValue(a int) {
	a = 10
}

// 比较并且交换
func swapValue(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

// 比较并且交换地址
func swapRefValue(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

	fmt.Println(temp)
}

func main() {

	a := 1
	changeValue(a)
	fmt.Println("a = ", a)
	changeRefValue(&a)
	fmt.Println("a = ", a)

	//比较并交换
	var (
		x = 1
		y = 2
	)

	swapValue(x, y)
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	swapRefValue(&x, &y)
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	//一级指针
	var p *int
	p = &x

	//二级指针
	var pp **int
	pp = &p

	fmt.Println("p=", p)
	fmt.Println("*p=", *p)
	fmt.Println("pp=", pp)
	fmt.Println("*pp=", *pp)
}
