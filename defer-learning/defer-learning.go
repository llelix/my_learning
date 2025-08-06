package main

import "fmt"

//defer延迟执行
//会在return 真正的R之前去执行
//多个defer采用栈 函数先进后出

func f(a int) int {

	defer func(a int) {
		a = a + 1
	}(a)
	a = a - 1

	return a
}

func main() {
	c := 3
	d := f(c)
	fmt.Println(d)

}
