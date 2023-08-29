package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := "gid://shopify/DiscountCodeNode/1151160320163"
	re := regexp.MustCompile(`/(\d+)$`)
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		finalNumberStr := match[1]
		finalNumber, err := strconv.ParseInt(finalNumberStr, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Final Number (int64):", finalNumber)
		}
	} else {
		fmt.Println("No match found.")
	}
}
