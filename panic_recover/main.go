package main

import "fmt"

func repairFunc() {
	if r := recover(); r != nil {
		fmt.Println("Function repaired")
	}
}

func studentIsApproved(n1, n2 float32) bool {
	defer repairFunc()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media is 6!")
}

func main() {
	fmt.Println(studentIsApproved(6, 6))
	fmt.Println("After execution") // with only panic, this should not be executed
}
