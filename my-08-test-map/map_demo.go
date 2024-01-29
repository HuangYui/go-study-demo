package main

import (
	"fmt"
)

func main() {
	//第一种方式
	var aMap map[string]string
	aMap = make(map[string]string)
	aMap["什么语法"] = "????"

	fmt.Println(aMap)

	//第二种
	bMap := make(map[string]string)
	bMap = make(map[string]string)
	bMap["什么语法"] = "????"

	fmt.Println(bMap)

	//第三种
	cmap := map[string]string{
		"a": "1",
		"b": "2",
	}
	fmt.Println(cmap)

}
