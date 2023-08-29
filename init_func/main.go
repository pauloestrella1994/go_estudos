package main

import "fmt"

// could have one for each file
func init() {
	fmt.Println("executing before main")
}

func main() {
	fmt.Println("executing on main")
}
