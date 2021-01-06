package main

import (
	"fmt"
	configChecker "github.com/airatbikbaev/geekbrains-go/configs"
)

func main() {
	isConfigCorrect := configChecker.IsConfigCorrect()

	if !isConfigCorrect {
		fmt.Println("Конфиг настроен неверно!")

		return
	}

	fmt.Println("Конфиг настроен верно!")
}
