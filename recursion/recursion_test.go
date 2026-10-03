package recursion

import "testing"

func TestFactorial(t *testing.T) {
	res := Factorial(5)

	expected := 120

	if res != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, res)
	}
}

func TestGCD(t *testing.T) {
	res := GCD(62, 8)

	expected := 2

	if res != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, res)
	}
}

func TestFindingExponent(t *testing.T) {
	res := FindingExponent(2, 4)

	expected := 16

	if res != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, res)
	}
}

func TestFibonacci(t *testing.T) {
	res := Fibonacci(10)

	expected := 55

	if res != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, res)
	}
}
