package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type PHPProgrammer struct {
}

func (p PHPProgrammer) Speak() string {
	return "PHP is the best language! "
}

type PassBy struct {
}

func main() {
	animals := []Animal{Dog{}, Cat{}, PHPProgrammer{}}
	// animals[3] = PassBy{}		// 编译不会通过, 因为没有实现 Animal 接口
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
