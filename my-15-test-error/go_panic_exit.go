package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		defer func() {
			a := recover()
			if a != nil {
				fmt.Println(a)
			}
		}()
		time.Sleep(2 * time.Second)
		panic("协程抛出了个异常")
	}()

	fmt.Println("hello!")
	time.Sleep(5 * time.Second)
	fmt.Println("嘿嘿嘿!")

}
