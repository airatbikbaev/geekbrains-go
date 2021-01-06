package main

import (
	"fmt"
	configChecker "github.com/airatbikbaev/geekbrains-go/configs"
	hw3 "github.com/airatbikbaev/geekbrains-go/hw3"
)

func main() {
	array := []int{20, 34, 0, 12, 5, 17, 10}
	sortedBubbleArray := hw3.GetBubbleSortedSlice(array)
	sortedIntersectionArray := hw3.GetInsertionSortedSlice(array)

	fmt.Println(array)
	fmt.Println(sortedBubbleArray)
	fmt.Println(sortedIntersectionArray)
}
