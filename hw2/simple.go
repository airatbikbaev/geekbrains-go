package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := scanNumber()

	printSimpleNumbers(int(n))
}

func scanNumber() int64 {
	var x string

	for {
		fmt.Println("Введите число N")
		fmt.Scanln(&x)

		val, err := strconv.ParseInt(x, 10, 64)

		if err != nil {
			continue
		} else {
			return val
		}
	}
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
