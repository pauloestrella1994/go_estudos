package main

import "fmt"

func main() {
	var array1 [5]string
	array1[0] = "first position"
	fmt.Println(array1)

	array2 := [3]string{"first position", "second position", "third position"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 54, 5, 6}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// slice works like a pointer to the array
	// by changing the position of the array, the slice changed
	array2[1] = "Position changed"
	fmt.Println(slice2)
}
