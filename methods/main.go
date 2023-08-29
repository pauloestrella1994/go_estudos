package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving data, user name: %s\n", u.name)
}

func (u user) checkAge() bool {
	return u.age >= 18
}

func (u *user) incrementingAge() {
	u.age++
}

func main() {
	user2 := user{"Paulo", 21}
	fmt.Println(user2)
	user2.save()
	fmt.Println(user2.checkAge())
	user2.incrementingAge()
	fmt.Println(user2.age)
}
