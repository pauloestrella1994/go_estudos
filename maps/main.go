package main

import "fmt"

func main() {
	user := map[string]string{
		"name":      "paul",
		"last_name": "netto",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {"first": "paul", "last": "netto"},
	}

	fmt.Println(user2["name"]["first"])

	delete(user2["name"], "last")

	fmt.Println(user2)

	user2["name"] = map[string]string{
		"age": "29",
	}

	fmt.Println(user2)
}
