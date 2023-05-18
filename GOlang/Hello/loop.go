package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {

		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			fmt.Println("DIVISIBLE By Both")

		} else if divisibleBy3 {
			fmt.Println("DIVISIBLE By 3")
		} else if divisibleBy5 {
			fmt.Println("DIVISIBLE By 3")
		} else {
			fmt.Println(i)
		}
	}

}
