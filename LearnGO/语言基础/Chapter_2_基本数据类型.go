/*
 * @Author: Firewine
 * @Date: 2020-04-30 10:05:01
 * @LastEditTime: 2020-04-30 16:35:13
 */

package main

import (
	"fmt"
	"math"
	"unicode"
	"unsafe"
)

/*
	一 整型
   1. 按照长度分为 int8，int16，int32，int64 对应的无符无符号整型：uint8，uint16，uint32，uint4
   2. uint8 为 byte型，int16 为c语言中的short型，int64为 long
*/

/*
	二 数字字面量语法
	可以使用二进制，八进制，十六进制浮点数的格式定义数字
	v: oboo101101
*/
func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 三 浮点型
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 四 复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 五 字符串
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	// 5.1 多行字符串
	si := `第一行
	第二行
	第三行`
	fmt.Println(si)
	/*
		字符串常用方法
			len(str) 	求长度
		+或fmt.Sprintf 	拼接字符串
		strings.Split 	分割
		strings.contains 	判断是否包含
		strings.HasPrefix,strings.HasSuffix 	前缀/后缀判断
		strings.Index(),strings.LastIndex() 	子串出现的位置
		strings.Join(a[]string, sep string) 	join操作
	*/

	const str = "hello沙河小王子"
	countStringCh(str)

	// 查看变量的字节大小和数据类型（使用较多  ）
	//	var n2 int64 = 10
	var n1 = 100
	// unsafe.sizeof(n1) 是返回变量占用的字节数
	fmt.Printf("n2 的类型 %T  n2 占用的字节数 %d ", n1, unsafe.Sizeof(n1))
}

// 六，在单个字符串为字符的形式，可以通过遍历和单个获取字符串，字符采用单引号包括起来
// 6.1 当需要处理中文、日文、或者其他复合字符，需要使用rune类型，实际也是一个int32
// 遍历字符串
//une类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}

	fmt.Println()
}
func countStringCh(string2 string) {
	hzc := 0
	for _, r := range string2 {
		if unicode.Is(unicode.Han, r) {
			hzc++
		}
	}
	fmt.Printf("汉字的数量：%v ", hzc)
}

//6.2 要修改字符串，需要先将其转换成rune 或者byte【】 ，完后，再转换成string，无
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 七 类型转换
//1. 没有java的 隐式转换，在go里面，只能显式强制类型转换
//t(表达式)
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
