package main

import "fmt"

func printfInterfaceArg(arg interface{}) {
	fmt.Println(arg)
	value, ok := arg.(string)
	fmt.Println(ok)
	fmt.Println(value)

}

func main() {
	printfInterfaceArg("aa")
	printfInterfaceArg(1)
}
