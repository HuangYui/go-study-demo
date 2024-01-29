package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	//数组截取（引用传递）
	s2 := slice1[0:2]
	//这里改了第一个元素，slice1 和 s2 的值都变了
	s2[0] = 100
	fmt.Println(slice1)
	fmt.Println(s2)

	slice3 := make([]int, len(slice1))

	//拷贝 不是拷贝内存地址，而是值传递
	copy(slice3, slice1)

	slice3[0] = 5555
	fmt.Println(slice3)

}
