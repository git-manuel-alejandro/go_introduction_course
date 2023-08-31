package funcs

import (
	"errors"
	"fmt"
)

func FactoryOperation(op Operation) func(float64, float64) float64 {

	switch op {
	case SUM:
		return sum
	case SUB:
		return sub
	case DIV:
		return div
	case MUL:
		return mul
	}

	return nil
}

func sub(a, b float64) float64 {
	return a - b
}

func sum(a, b float64) float64 {
	return a + b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func Suma(values ...float64) float64 {
	var sum float64
	fmt.Println(values)
	for _, v := range values {
		sum += v
	}
	return sum
}

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil

	case SUB:
		return x - y, nil

	case DIV:
		if y == 0 {
			return 0, errors.New("y mustn't be zero")
		}
		return x / y, nil

	case MUL:
		return x * y, nil
	}

	return 0, errors.New("Something was wrong")

}

func PrintString(n int, text string) {
	for i := 0; i < n; i++ {
		fmt.Println(text)
	}
}

func Add(x, y int) int {
	return x + y
}
