package main

import (
	"../utils"
	"fmt"
)

func main() {
	var a, b float64

	a = utils.ScanFloatNumber("Введите число А:")
	b = utils.ScanFloatNumber("Введите число B:")

	fmt.Printf("Результат вычислений: %f\n", getCalculation(a, b))
}

func getCalculation(a float64, b float64) float64 {
	var operation string

	for {
		fmt.Println("Введите операцию:")
		fmt.Scanln(&operation)

		switch operation {
		case "+":
			return a + b
		case "-":
			return a - b
		case "/":
			return a / b
		case "*":
			return a * b
		case "%":
			return float64(int(a) % int(b))
		case "^":
			x := a

			for i := 1; i < int(b); i++ {
				x *= a
			}

			return x
		default:
			continue
		}
	}
}
