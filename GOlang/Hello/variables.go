package main

import "fmt"

func main() {
	// storing favourite color in a variable using "var keyword"
	var favcolor = "red"
	fmt.Println("my favourite color is", favcolor)

	// storing two values in two variables using compound assignment
	birthYear, ageInYears := 2000, 22
	fmt.Println("Born in", birthYear, "aged", ageInYears)

	// storing two values in two variables using block assignment
	var (
		first  = 'J'
		second = 'L'
	)
	fmt.Println("First", first, "Second", second)

	// declaring but not assigning
	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("I am", ageInDays, "days old")

}
