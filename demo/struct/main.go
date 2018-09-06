package main

import (
	"fmt"
)

type human struct {
	Name string
	Age  int
}

// 隐式的继承 human 的属性
type ming struct {
	human
	Skill string
}

// 添加结构体所属的方法
func (h human) Say() {
	fmt.Printf("我叫%s,今年%d", h.Name, h.Age)
}

func main() {
	var xm ming
	xm.Name = "小明"
	xm.Age = 18
	xm.Skill = "跑得快"
	fmt.Println(xm)

	xm.Say()
}
