
package main

import "fmt"

/*
1. 四 流程控制

*/

func main() {
	// 1. if else 分支结构
	//
	if false {
		fmt.Print("分支1")
	}else{
		fmt.Print("分支2")
	}
	// go 语言规定与 if匹配的左括号必须与if 和表达式放在同一行，放在其他地方会报错



}
func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
// 1.1 if条件的判断的特殊语法
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// 2.  for 结构
// 在go 语言中，所有循环类型均可以使用for关键字来完成，没有while循环
// 没有java 那么规范化，初始语句和结束语句都是可以省略
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 2.2 go语言没有while循环，则使用for来替代
	for {
		fmt.Println("ceshi")
		break
		// 使用break，goto,return，panic 都是可以强制退出循环
	}
//	2.3 可以使用for range 遍历数组，切片，字符串，map，及通道，
/*

   数组、切片、字符串返回索引和值。
   map返回键和值。
   通道（channel）只返回通道内的值。

*/


}
// 3.  switch case
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
// 3.1 一个分支可以有多个值，多个case 值中间使用英文逗号分隔
func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
//	3.2 使用表达式
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

//	3.3 fallthrough 语法可以执行满足条件case的下一个case，兼容c语言的case设计
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
// 4 goto 跳转到指定标签, 以前是说goto会破坏 代码的结构性，注意下
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}