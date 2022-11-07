package main

import "fmt"

// var 语句可以出现在包或函数级别
var aa int64

func main() {
	// var 语句用于声明一个变量列表
	var a int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b float32 // float64
	var c bool
	var d string
	var e byte // 等同于 uint8
	var f rune // 等同于 int32,表示一个 Unicode 码点
	var g interface{}

	// 没有明确初始值的变量声明会被赋予它们的 零值
	// 零值是:
	//    数值类型为 0
	//    布尔类型为 false
	//    字符串为 ""（空字符串）

	fmt.Println("int zero value: ", a)
	fmt.Println("int64 zero value: ", aa)
	fmt.Println("float32 zero value: ", b)
	fmt.Println("bool zero value: ", c)
	fmt.Println("string zero value: ", d)
	fmt.Println("byte zero value: ", e)
	fmt.Println("rune zero value: ", f)
	fmt.Println("rune zero value: ", g)
}
