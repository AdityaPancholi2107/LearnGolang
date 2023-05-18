package calculator_test

import (
	"testing"
	"maths/math_fun"
)

func TestMultiply(t *testing.T) {
	result := calculator.Multiply(5, 4)
	expected := 20

	if result != expected {
		t.Errorf("Multiply failed: expected %d, got %d ", expected, result)
	}
}
func TestDivide(t *testing.T) {
	result := calculator.Divide(24.00, 4.00)
	expected := 6.00

	if result != expected {
		t.Errorf("Divide failed: expected %f, got %f ", expected, result)
	}
}
