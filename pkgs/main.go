package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing in main file")
	assistant.Write()
	err := checkmail.ValidateFormat("foo@gmail.com")
	fmt.Println(err)
}
