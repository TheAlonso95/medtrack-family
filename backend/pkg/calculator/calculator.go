package calculator

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
// Panics if b is zero
func Divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}