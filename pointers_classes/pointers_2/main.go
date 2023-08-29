package main

import "fmt"

func Inverted(num int) int {
	return num * -1
}

func InvertedNum(num *int) {
	*num = *num * -1 // no need return, modifying the variable directly
}

func main() {
	num := 20
	invertedNum := Inverted(num)
	fmt.Println(invertedNum)
	fmt.Println("num not modified:", num)

	newNum := 40
	fmt.Println(newNum)
	InvertedNum(&newNum)
	fmt.Println(newNum)
}
