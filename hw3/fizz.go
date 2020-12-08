package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		mod3 := i%3 == 0
		mod5 := i%5 == 0

		if mod3 && mod5 {
			fmt.Println("FizzBuzz")

			continue
		}

		if mod3 {
			fmt.Println("Fizz")

			continue
		}

		if mod5 {
			fmt.Println("Buzz")

			continue
		}

		fmt.Println(i)
	}
}
