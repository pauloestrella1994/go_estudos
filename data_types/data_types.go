package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = -10000000000
	fmt.Printf("%v, %T\n", number, number)

	var number1 uint64 = 1000
	fmt.Printf("%v, %T\n", number1, number1)

	// rune alias to i32
	var number2 rune = -1000000000
	fmt.Printf("%v, %T\n", number2, number2)

	// byte alias to unit8
	var number3 byte = 10
	fmt.Printf("%v, %T\n", number3, number3)

	var number4 float32 = 10.0
	fmt.Printf("%v, %T\n", number4, number4)

	number5 := 10.0
	fmt.Printf("%v, %T\n", number5, number5)

	var str string = "Text"
	fmt.Printf("%v, %T\n", str, str)

	// return ascii value
	char := 'A'
	fmt.Printf("%v, %T\n", char, char)

	// initial values
	var text string
	fmt.Printf("%v, %T\n", text, text)

	var num int32
	fmt.Printf("%v, %T\n", num, num)

	var boolean bool
	fmt.Printf("%v, %T\n", boolean, boolean)

	var err error
	fmt.Printf("%v, %T\n", err, err)

	err2 := errors.New("Error object")
	fmt.Printf("%v, %T\n", err2, err2)
}
