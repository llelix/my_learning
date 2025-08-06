package main

import "fmt"

//定义结构体
type Hero struct {
	//定义成员变量
	Name  string
	Level int
}

//成员方法
//接收者为pointer
func (hero *Hero) setNameP() {
	hero.Level = 100
}

//接收者为组类型
func (hero Hero) setName() {
	hero.Level = 100
}

//结构体的内存对齐

//结构体
func main() {

	var hero Hero
	fmt.Printf("ss%v \n", hero)
	hero.setName()
	fmt.Printf("ss%v \n", hero)
	hero.setNameP()
	fmt.Printf("ss%v \n", hero)
	//解构赋值
	hero2 := Hero{"ss", 100}
	//赋值
	hero3 := Hero{Name: "是", Level: 101}
	fmt.Println(hero2, hero3)

}
