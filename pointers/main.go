package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var2++
	fmt.Println(var1, var2)

	// pointer is a reference to a memory

	var var3 int64
	var pointer *int64

	fmt.Println(var3, pointer)

	var3 = 100
	pointer = &var3

	var3++
	fmt.Println(&var3, pointer) // memory address
	fmt.Println(var3, *pointer) // dereference
}
