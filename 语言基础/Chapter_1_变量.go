package main

import "fmt"

// 5. 全局变量
var m = 100

// 函数
func foo() (int, string) {
	return 10, "Q1mi"
}
func main() {

	//1.  go 语言的变量声明格式为
	// var 变量名 变量类型
	var name string = "demo"
	//2. go 语言的批量声明
	var (
		a string
		b int
		c bool
	)
	//3. 变量的初始化
	var name1 string = "qimi"
	//4. 可以不用进行类型显式声明， 可以交给程序自己判断
	var name2 = "q1mi"
	// 5. 此处声明局部变量
	m := 200

	// 声明匿名变量，在使用多重赋值，想要忽略某个值，可以使用匿名变量用_表示
	x, _ := foo()
	fmt.Println("x=", x)
	// 匿名斌良，不占用命名空间，不会分配内存
	/*
			 注意事项：
		    	函数外的每个语句都必须以关键字开始（var、const、func等）
		    	:=不能使用在函数外。
		    	_多用于占位，表示忽略值。

	*/
	// 6. 常量 只是将var 变为const,也可以支持多个常量同时声明
	const pi = 3.1415

	// 7. iota 是go语言的常量计数器，只能在常量的表达式中使用
	// 在 const关键词出现时被重置为0，const每新增一行常量声明将使iota计数一次
	const (
		n1 = iota //0
		n2        // 1
		n3        //2
		n4        //3
		//定义枚举非常的适用
	)
	// 7.1 补充，使用下划线可以 跳过某些值
	const (
		na = iota
		nb //1
		_
		nc //
	)
	//7.5 可以在批量插入声明下，可以插入值
	const n5 = iota // 0
	fmt.Println(m, name2, name1, name, a, b, c)
	fmt.Println("hello")
}
