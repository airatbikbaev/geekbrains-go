package main

import (
	"../utils"
	"errors"
	"fmt"
	"math"
)

func main() {
	var a, b float64

	a = utils.ScanFloatNumber("Введите число А:")
	b = utils.ScanFloatNumber("Введите число B:")

	calc, err := getCalculation(a, b)

	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)

		return
	}

	fmt.Printf("Результат вычислений: %f\n", calc)
}

func getCalculation(a float64, b float64) (float64, error) {
	var operation string

	for {
		fmt.Println("Введите операцию:")
		fmt.Scanln(&operation)

		switch operation {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		case "/":
			if b == 0 {
				return 0, errors.New("На 0 делить нельзя")
			}
			return a / b, nil
		case "*":
			return a * b, nil
		case "%":
			return math.Mod(a, b), nil
		case "^":
			x := a

			for i := 1; i < int(b); i++ {
				x *= a
			}

			if b < 0 {
				return 1 / x, nil
			}

			return x, nil
		default:
			continue
		}
	}
}
