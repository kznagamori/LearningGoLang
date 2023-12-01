package main

import (
	"fmt"
)

// shellSort はシェルソートを実行する関数です。
func shellSort(arr []int) []int {
	n := len(arr)

	// ギャップを縮小しながら挿入ソートを行う。
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for ; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = temp
		}
	}

	return arr
}

func main() {
	arr := []int{12, 34, 54, 2, 3}
	fmt.Println("Original array:", arr)
	sortedArr := shellSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
