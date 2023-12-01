package main

import (
	"fmt"
)

// quickSort はクイックソートを実行する関数です。
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// ピボットは中間の要素とします。
	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{10, 5, 2, 3, 7, 6, 1, 9, 8, 4}
	fmt.Println("Original array:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
