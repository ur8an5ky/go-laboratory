package operations

func Factorial(num1 int) int {
	if num1 == 1 {
		return 1
	}
	return num1 * Factorial(num1-1)
}
