package utils

import (
	"fmt"
	"strconv"
)

func ScanFloatNumber(text string) float64 {
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

func ScanIntNumber(text string) int64 {
	var x string

	for {
		fmt.Println(text)
		fmt.Scanln(&x)

		val, err := strconv.ParseInt(x, 10, 64)

		if err != nil {
			continue
		} else {
			return val
		}
	}
}

func ScanUintNumber(text string) uint64 {
	var x string

	for {
		fmt.Printf(text)
		fmt.Scanln(&x)

		val, err := strconv.ParseUint(x, 10, 64)

		if err != nil {
			continue
		} else {
			return val
		}
	}
}
