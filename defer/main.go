package main

import "fmt"

func func1() {
	fmt.Println("Executing first function...")
}

func func2() {
	fmt.Println("Executing second function...")
}

func studentIsApproved(n1, n2 float32) bool {
	defer fmt.Println("The result is: ")

	fmt.Println("Entering on the function to see if student is approved")

	media := (n1 + n2) / 2

	if media >= 5 {
		return true
	}

	return false
}

func main() {
	defer func1()
	func2()

	fmt.Println(studentIsApproved(7, 8))
}
