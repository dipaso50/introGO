package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Move() {
	fmt.Printf("%s is Walking\n", a.Name)
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

type Fish struct {
	Animal
}

func (a Fish) Move() {
	fmt.Printf("%s is swimming\n", a.Name)
}

func main() {
	d := Dog{Animal{Name: "firulais"}}
	d.Move()

	j := Dog{Animal{Name: "Jhon"}}
	j.Move()

	f := Fish{Animal{Name: "Nemo"}}
	f.Move()
}
