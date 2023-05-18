package main

import "fmt"

type Calculator struct {
	value_1  float64
	value_2  float64
	operator string
}

func (c Calculator) add() float64 {
	return c.value_1 + c.value_2
}
func (c Calculator) subtract() float64 {
	return c.value_1 - c.value_2
}
func (c Calculator) multiply() float64 {
	return c.value_1 * c.value_2
}
func (c Calculator) divide() float64 {
	return c.value_1 / c.value_2
}
func (c Calculator) calculte() (float64, error) {
	switch c.operator {
	case "+":
		return c.add(), nil
	case "-":
		return c.subtract(), nil
	case "*":
		return c.multiply(), nil
	case "/":
		return c.divide(), nil
	default:
		return 0, fmt.Errorf("Unknown Operator")
	}

}
func main() {
	var c Calculator

	fmt.Println("Enter First Number:")
	fmt.Scanln(&c.value_1)

	fmt.Println("Enter Operator:")
	fmt.Scanln(&c.operator)

	fmt.Println("Enter Second Number:")
	fmt.Scanln(&c.value_2)

	result, _ := c.calculte()
	//if err != nil {
		//panic(err)
	

	fmt.Printf("%.2f %s %.2f = %.2f", c.value_1, c.operator, c.value_2, result)

}
