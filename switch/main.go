package main

import "fmt"

func day_of_week(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Out of the range"
	}
}

func day_of_week2(number int) string {
	var dow string

	switch {
	case number == 1:
		dow = "Sunday"

		fallthrough
	case number == 2:
		dow = "Monday"

		fallthrough
	case number == 3:
		dow = "Tuesday"

		fallthrough
	case number == 4:
		dow = "Wednesday"

		fallthrough
	case number == 5:
		dow = "Thursday"

		fallthrough
	case number == 6:
		dow = "Friday"

		fallthrough
	case number == 7:
		dow = "Saturday"

		fallthrough
	default:
		dow = "Out of the range"

	}

	// fallthrough is a statement to jump your clause to another conditional
	return dow
}

func main() {
	fmt.Println("Switch")

	day := day_of_week(3)
	fmt.Println(day)

	fmt.Println("-------------")

	day2 := day_of_week2(3)
	fmt.Println(day2)

	fmt.Println("-------------")

	day3 := day_of_week(8)
	fmt.Println(day3)
}
