package main

import "fmt"

func closure() func() {
	text := "Inside of closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Inside of main function"
	fmt.Println(text)

	newFunction := closure() // closure function will maintain the text variable reference from closure function
	newFunction()
}
