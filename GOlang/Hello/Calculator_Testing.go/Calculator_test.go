package calculator

import(
	"testing"
	

)

func TestAdd(t*testing.T){
	result := Add(2.00, 3.00)
	expected := 5.00
	
	if result != expected{
		t.Errorf("Add Function failed: expected %f, got %f", expected, result)
	}
}

func TestSubtract(t*testing.T){
	result := Subtract(3.00, 2.00)
	expected := 1.00
	if result != expected{
		t.Errorf("Subtract function failed: expected %f, got %f", expected, result)
	}
}