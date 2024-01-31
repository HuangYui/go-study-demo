package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子go程结束")

		fmt.Println("子go程正在运行……")

		c <- 666 //666发送到c
		fmt.Println("发送channel后子go程正在运行……")

	}()
	time.Sleep(3 * time.Second)
	num := <-c //从c中接收数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main go程结束")

}
