package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Addr string
}

func (receiver User) GetName() (a, b int) {
	fmt.Println("user 执行了....")
	return 1, 2
}

func main() {
	var a int
	of := reflect.TypeOf(a)
	valueOf := reflect.ValueOf(a)
	fmt.Println(of)
	fmt.Println(valueOf)

	var user = User{"张三", "广州"}
	typeOf := reflect.TypeOf(user)
	value := reflect.ValueOf(user)
	//字段的数量
	field := typeOf.NumField()
	for i := 0; i < field; i++ {
		//字段
		structField := typeOf.Field(i)
		//名称
		println(structField.Name)
		v := value.Field(i)
		//值
		fmt.Println(v.Interface())
	}

	//获取接口
	method := typeOf.NumMethod()
	for i := 0; i < method; i++ {
		m := typeOf.Method(i)
		fmt.Println(m.Name)
	}
}
