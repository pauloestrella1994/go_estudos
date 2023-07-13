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

	// internal arrays
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	// when you exceed the array size, the size duplicate to carry all the values
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)

	// when you exceed the array size, the size duplicate to carry all the values
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
