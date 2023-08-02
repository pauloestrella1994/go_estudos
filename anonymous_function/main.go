package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello World!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Hello World!!")

	data := func(text string) string {
		return fmt.Sprintf("Hello %s", text)
	}("World!!!")

	fmt.Println(data)
}
