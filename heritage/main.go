package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type student struct {
	person
	college string
}

func main() {
	p1 := person{"john", 10}

	fmt.Println(p1)

	e1 := student{p1, "UCLA"}
	fmt.Println(e1.name)
	fmt.Println(e1.age)
	fmt.Println(e1.college)
}
