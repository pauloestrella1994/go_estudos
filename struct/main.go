package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
}

func main() {
	var u user

	u.age = 8
	u.name = "Paulo"
	fmt.Println(u)

	u2 := user{"Brad", 10, address{"foo"}}
	fmt.Println(u2)

	u3 := user{name: "Brad"}
	fmt.Println(u3)

}
