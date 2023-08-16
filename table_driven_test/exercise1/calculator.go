package exercise1

import "fmt"

const (
	ErrInvalidNumber = "invalid number"
	ErrDivideByZero  = "cannot divide by zero"
)

func Add(first, second int) int {
	return first + second
}

func Subtract(first, second int) int {
	return first - second
}

func Multiplication(first, second int) int {
	return first * second
}

func Division(first, second int) (float32, error) {
	if second == 0 {
		return float32(0), fmt.Errorf(ErrDivideByZero)
	}
	return float32(first) / float32(second), nil
}

func Factorial(num int) (int, error) {
	if num < 0 {
		return 0, fmt.Errorf(ErrInvalidNumber)
	}

	var result int = 1

	for i := 1; i <= num; i++ {
		result = i * result
	}

	return result, nil
}
