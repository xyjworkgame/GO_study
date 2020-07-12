package main

import "fmt"

// 5 数组array

func main() {
//	5.1 定义形式
	var a [3]int;
	var b =[3]int{1,2,3}
	// 指定的时候，不需要全部容量都必须初始化
	var c = [3]int{1,2}// 不够是以0为默认值
	var d = [...]int{1,2}
	// 让编辑器根据初始化值，自己判断容量大小
	// 指定索引值来进行初始化
	var e = [...]int{1:1,3:5}

	fmt.Print(a,b,c,d,e)
}

// 5.2 数组的遍历方法
func sys(){
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
// 5.3 二分数组
func twoArray(){
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Print(a)
}
