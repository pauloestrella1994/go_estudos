package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 5 {
		time.Sleep(time.Second)
		fmt.Printf("Counting %v \n", i)

		i++
	}

	fmt.Println("-------------")

	for j := 0; j < 3; j++ {
		time.Sleep(time.Second)
		fmt.Printf("Counting %v \n", j)
	}

	fmt.Println("-------------")

	names := [3]string{"john", "paul", "vince"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	fmt.Println("-------------")

	for index, char := range "WORD" {
		fmt.Println(index, string(char))
	}

	fmt.Println("-------------")

	user := map[string]string{
		"name": "Paul",
		"age":  "29",
	}

	for index, value := range user {
		fmt.Println(index, string(value))
	}

	// cannot iterate over structs
	// for {
	// 		infinite loop
	// }
}
