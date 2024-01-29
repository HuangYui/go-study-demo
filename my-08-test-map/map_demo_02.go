package main

import "fmt"

func main() {

	aMap := map[string]string{
		"hello": "world",
		"ccc":   "www",
	}
	//map遍历

	for key, value := range aMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	//map新增
	aMap["what"] = "!!!!!"

	for key, value := range aMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	//map删除
	delete(aMap, "what")
	for key, value := range aMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}
