package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func sumAndSubtraction(n1, n2 int8) (int8, int8) {
	return n1 + n2, n1 - n2
}

func main() {
	sum := sum(10, 20)
	fmt.Println(sum)

	var f = func(txt string) string {
		fmt.Println(txt)
		return (txt)
	}

	f("H function")
	result := f("F function")
	fmt.Printf(result)

	sumResult, subsResult := sumAndSubtraction(8, 4)
	fmt.Printf("\nSum: %v\nSubtraction: %v\n", sumResult, subsResult)
}
