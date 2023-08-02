package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}

func write(text string, numbers ...int) {
	total := 0

	for _, n := range numbers {
		total += n
	}

	fmt.Println(text + fmt.Sprint(total))

}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	write("Total: ", 1, 2, 3, 4, 5)
}
