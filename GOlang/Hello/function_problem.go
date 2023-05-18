package main

import "fmt"

func twoNumbers() (int, int) {
	return 5, 5
}
func anyNumber() int {
	return 10
}
func addingThreeNo(a, b, c int) int {
	return a + b + c
}
func greetingPerson(name string) {
	fmt.Println("Hello", name)
}
func hiThere() string {
	return "Hi There!"
}

func main() {
	greetingPerson("AP")
	fmt.Println(hiThere())
	a, b := twoNumbers()
	answer := addingThreeNo(anyNumber(), a, b)
	fmt.Println(answer)

}
