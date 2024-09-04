package calculator

type Calculator struct {
}

func CreateCalculator() Calculator {
	return Calculator{}
}
func (c Calculator) add(a int, b int) int {
	return a + b
}

func (c Calculator) subtract(a int, b int) int {
	return a - b
}
func (c Calculator) multiply(a int, b int) int {
	return a * b
}
func (c Calculator) divide(a int, b int) int {
	return a / b
}

func (c Calculator) MultiplyNumbers(numbers []int) int {
	result := 1

	for _, value := range numbers {
		result *= value
	}
	return result
}
