package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) CelebrateBirthday() {
	p.age++
}

func main() {
	p1 := Person{}

	p1.name = "Jhon"
	p1.age = 30

	fmt.Printf("%v \n", p1)

	p2 := Person{name: "Jane", age: 25}

	fmt.Printf("%v \n", p2)

	p2.CelebrateBirthday()

	fmt.Printf("%v \n", p2)
}
