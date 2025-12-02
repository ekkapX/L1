package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left := -1
	right := len(arr)
	for right-left > 1 {
		mid := (left + right) / 2
		if arr[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	if right < len(arr) && arr[right] == target {
		return right
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 1))
}
