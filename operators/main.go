package main

import "fmt"

func main() {
	sum := 1 + 2
	sub := 2 - 1
	div := 10 / 4
	mult := 10 * 5
	rest := 10 % 2

	fmt.Println(sum, sub, div, mult, rest)

	var var1 string = "String"
	var2 := "String2"
	fmt.Println(var1, var2)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	ver, fal := true, false
	fmt.Println(ver && fal)
	fmt.Println(ver || fal)
	fmt.Println(!ver)

	num := 10
	num++
	num += 15
	num *= 3
	num /= 3
	num %= 3

	fmt.Println(num)

}
