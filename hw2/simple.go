package main

import (
	"../utils"
	"fmt"
)

func main() {
	n := utils.ScanIntNumber("Введите число N:")

	printSimpleNumbers(int(n))
}

func printSimpleNumbers(n int) {
	fmt.Println("Простые числа: ")

	operationsCount := 0

	for i := 1; i <= n; i++ {
		condition := false

		for j := 2; j < i; j++ {
			operationsCount += 1

			if i%j == 0 {
				condition = true

				break
			}
		}

		if !condition {
			fmt.Println(i)
		}
	}

	fmt.Printf("Кол-во операций: %d\n", operationsCount)
}
