package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b float64

	a = scanNumber("Введите число А:")
	b = scanNumber("Введите число B:")

	fmt.Printf("Результат вычислений: %f\n", getCalculation(a, b))
}

func scanNumber(text string) float64 {
	var x string

	for {
		fmt.Println(text)
		fmt.Scanln(&x)

		val, err := strconv.ParseFloat(x, 64)

		if err != nil {
			continue
		} else {
			return val
		}
	}
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
