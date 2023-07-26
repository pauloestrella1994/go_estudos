package main

import "fmt"

func main() {
	number := 10

	if number >= 15 {
		fmt.Println("Bigger or equal to 15")
	} else {
		fmt.Println("Less than 15")
	}

	if number2 := number; number2 > 9 {
		fmt.Println(number2, "bigger than 9")
	} else if number2 > 0 && number2 < 9 {
		fmt.Println(number2, "less than 9, bigger than zero")
	} else {
		fmt.Println(number2, "is negative")
	}
}
