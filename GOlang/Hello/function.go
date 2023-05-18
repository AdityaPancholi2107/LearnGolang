package main

import "fmt"

func double(x int) int {
	return x + x
}
func add(lhs, rhs int) int {
	return lhs + rhs
}
func intro() {
	fmt.Println("Hello !!!!")
}

func main() {
	intro()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("Bakers Dozen is", bakersDozen)

	another := add(double(6), 1)
	fmt.Println("Another Value", another)
}
