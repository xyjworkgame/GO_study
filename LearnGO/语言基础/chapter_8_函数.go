/*
@Time : 2020/5/1 18:26
@Author : Firewine
@File : chapter_8_函数
@Software: GoLand
*/
package main

import "fmt"

/*
	go 语言支持 函数、匿名函数和闭包
格式
func 函数名(参数)(返回值){函数体}
多个变量和多个返回值用逗号隔开
若是多个参数 类型相同，可以省略后面的
*/

//8.2 可变参数的描述
func intSum2(x ...int)int{
	fmt.Println(x) // x是一个切片
	sum := 0
	for _,v := range x{
		sum += sum +v
	}
	return sum
}
// 8.3 return 返回值可以提前指定，并且在函数体使用
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// nil 可以说是 null，或者 是 0 作为返回值，切片则为是nil返回


// 8.4 函数类型与变量
// 8.4.1 定义函数类型
type  calculation func(int ,int )int
//它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
// 上面的函数直接都是calculation 类型的函数，
//add和sub都能赋值给calculation类型的变量。
//var c calculation
//c = add
func main6() {
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}
func main() {
	result := intSum2(1,2,3, 4)

	fmt.Println(result)
}