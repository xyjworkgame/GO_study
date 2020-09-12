/*
@Time : 2020/5/1 16:33
@Author : Firewine
@File : chapter_7_map
@Software: GoLand
*/
package main

import (
	"fmt"
	"strings"
)

/*
	七 map 是引用类型，必须初始化才能够使用
*/

// map[keyType]ValueType
// 使用make来分配内存
//make(map[keyType]valueType,[cap])
func main3() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
}

// 7.2 在使用过程中，判断map中键是否存在的特殊写法
func main2() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
// 删除元素 有方法delete


// 7.4 值为切片类型的map
func main1() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

// 统计一个字符串中字符串每个单词出现的次数
func main()  {

	str := "how do you do think you about"
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)

	countMap := make(map[string]int, 10)
	for _, key := range strSlice {
		_, isReal := countMap[key]
		if !isReal {
			countMap[key] = 1
		} else {
			countMap[key] += 1
		}
	}
	fmt.Println(countMap)
}
func countSentence(testString string) map[string]int {
	stringAfterop := strings.Split(testString," ")
	mapDemo := make(map[string]int,10)
	for _,v := range stringAfterop{
		mapDemo[v]++
	}
	return mapDemo
}