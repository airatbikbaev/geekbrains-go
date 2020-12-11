package main

import (
	"../utils"
	"fmt"
)

var fibonacciMap = map[uint64]uint64{0: 0, 1: 1, 2: 1}

func main() {
	n := utils.ScanUintNumber("Введите номер N:")

	fmt.Printf("Число Фибоначчи: %d\n", getFibonacciNumber(n))
}

// Сохраняем каждое посчитанное значение в мапу, чтобы при повторных обращениях доставать значение из памяти.
func getFibonacciNumber(n uint64) uint64 {
	x := fibonacciMap[n]

	if n <= 2 {
		return x
	}

	if x == 0 {
		fibonacciMap[n] = getFibonacciNumber(n-2) + getFibonacciNumber(n-1)

		return fibonacciMap[n]
	}

	return x
}
