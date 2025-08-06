package main

import "fmt"

//defer延迟执行
//会在return 真正的R之前去执行
//多个defer采用栈 函数先进后出

func f(a int) int { //区别在于在函数定义时定义的result变量
	var result int
	defer func(a int) {
		result += 1
	}(a)
	result = a - 1

	return result //到了return的时候result把值给R(真正返回值值)
}

func f2(a int) (result int) {

	defer func(a int) {
		result += 1
	}(a)
	result = a - 1

	return result
}

func main() {
	c := 3
	d := f(c)
	fmt.Println(d)
	e := f2(c)
	fmt.Println(e)

}
