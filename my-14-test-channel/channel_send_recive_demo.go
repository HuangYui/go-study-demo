package main

import (
	"fmt"
	"time"
)

// 双向 可发可收
func main() {
	c := make(chan int)

	go reciveAndSend(c)
	//主线程发送
	time.Sleep(3 * time.Second)
	c <- 2
	//主线程接收
	fmt.Println("主线程接收到了值", <-c)

}

func reciveAndSend(a chan int) {
	i := <-a
	fmt.Println("方法接收到了值", i)
	i++
	a <- i
}
