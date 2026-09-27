package recursion

func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func GCD(a, b int) int {
	m := a % b

	if m == 0 {
		return b
	}

	return GCD(b, m)
}

func FindingExponent(x, n int) int {
	if n == 0 {
		return 1
	}

	return x * FindingExponent(x, n-1)
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
