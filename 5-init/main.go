package main

import (
	_ "GolangStudy/5-init/lib1" //Go强制引入必须使用，如过需要引入了不使用则需要加_(会执行导入的init方法)
	"GolangStudy/5-init/lib2"
	mylib2 "GolangStudy/5-init/lib2"
	//. "GolangStudy/5-init/lib2"
)

func main() {
	//lib1.lib1Test()

	//lib2.Lib2Test()
	mylib2.Lib2Test()
	lib2.Lib2Test()
	//Lib2Test()
}
