package main

import "fmt"

type Jumper interface {
	Jump()
}

type Person struct {
	name string
}

type Animal struct {
	name string
}

func (p Person) Jump() {
	fmt.Printf("%s jumps!!\n", p.name)
}

func (p Animal) Jump() {
	fmt.Printf("%s jumps!!\n", p.name)
}

func main() {
	jane := Person{name: "Jane"}

	dog := Animal{"firulais"}

	doJump(jane)
	doJump(dog)
}

func doJump(jp Jumper) {
	jp.Jump()
}
