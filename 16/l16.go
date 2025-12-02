package main

import "fmt"

func patrition(arr []int, left, right int) int {
	pivot := arr[(left+right)/2]
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}

	}
	return left
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		index := patrition(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index, right)
	}
	return arr
}

func sort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
func main() {
	fmt.Println(sort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(sort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(sort([]int{1, 3, 1, 7, 4, 1, 5, 1, 8, 1}))
}
