package main

import "fmt"

// panic 运行中的异常，不捕获的话，会导致程序直接退出
func main() {
	//注册捕获panic的函数,必须先注册，若在panic之后则无效
	defer doPanic()
	n := 0
	res := 1 / n
	fmt.Println(res) //panic 之后的代码不会执行
}

// 当捕获到panic时触发此函数
func doPanic() {
	err := recover()
	if err != nil {
		fmt.Println("捕获到panic")
		fmt.Println(err)
	}
}
