package main

import "fmt"

/**
  6 切片，
  	是一个拥有相同类型元素的可变长度的序列，他是基于数组类型做的一层封装，且支持自动扩容。
	切片是一个引用类型，内部结构包含-地址-长度-容量，
	定义 var name []T
 */
func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
	var a1 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a1 = append(a1, fmt.Sprintf("%v", i))
	}
	fmt.Println(a1)
}
// 6.1 使用len 求长度，cap 求容量

//6.2 a[low : high : max] 完整的切片的表达是，a不能是字符串
func slice() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}
//t:[2 3] len(t):2 cap(t):4
// 完整切片表达式需要满足的条件是 0 <= low <= high <= max <= cap(a)


// 6.3 使用make 函数构造切片
//make([]T,size,cap) T 切片的元素类型，size 切片中的元素的数量，cap 切片的容量
func makeSlice() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}
// 6.4 切片的本质 ： 就是对底层数组的封装，他包含了三个信息，底层数组的指针，切片你的长度，和切片的容量
// 切片是否为空通过 长度是否为0来判断，并且切片不能直接比较
//使用append 方法进行去添加元素，扩容是 每次扩容前的2倍
//查看扩容源码 src/runtime/slice.go
// copy函数可以 复制切片，但是切片是引用类型，变量都是指向同一内存地址，修改同时发生变化
// 切片没有专门删除元素的方法，通过append 放通过元素index 进行删除
// a = append(a[:index],a[index+1]...)
