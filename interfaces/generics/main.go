package main

import "fmt"

func generics(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generics("String")
	generics(1)
	generics(true)

	// example of method that accepts a lot of interfaces types
	fmt.Println(1, 2, "String", true, float64(1234))

	maps := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(maps)
}
