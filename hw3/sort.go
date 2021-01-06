package hw3

func GetBubbleSortedSlice(arr []int) []int {
	sortedArray := make([]int, len(arr))

	copy(sortedArray, arr)

	for i := 0; i < len(sortedArray)-1; i++ {
		sorted := true

		for j := 0; j < len(sortedArray)-1; j++ {
			if sortedArray[j] > sortedArray[j+1] {
				tmp := sortedArray[j]
				sortedArray[j] = sortedArray[j+1]
				sortedArray[j+1] = tmp

				sorted = false
			}
		}

		if sorted {
			break
		}
	}

	return sortedArray
}

func GetInsertionSortedSlice(arr []int) []int {
	sortedArray := make([]int, len(arr))

	copy(sortedArray, arr)

	for i := 1; i < len(sortedArray); i++ {
		x := sortedArray[i]

		for j := i - 1; j >= 0; j-- {
			if x < sortedArray[j] {
				sortedArray[j+1] = sortedArray[j]
				sortedArray[j] = x
			} else {
				break
			}
		}
	}

	return sortedArray
}
