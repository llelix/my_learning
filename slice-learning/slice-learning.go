package main

import "fmt"

//slice就是对数组的切片
//有三个成员组成 pointer len cap的结构体

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	s1 := array[1:3] //左开右闭
	fmt.Println(s1)
	s2 := array[2:]
	fmt.Println(s2)
	//修改值会改变共享底层数组的其他切片
	s1[1] = 55
	fmt.Printf("s2 %v/n", s2)

	//容量不够的时候会扩容 翻倍  超过1024增加1/4
	//创建新的底层数组把切片拷贝过去

}
