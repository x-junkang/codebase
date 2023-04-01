package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := make([]int, 0)
	right := make([]int, 0)
	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = append(quickSort(left), pivot)
	return append(left, quickSort(right)...)
}

func quickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	tmp := arr[0]

	left, right := 0, len(arr)-1
	for left < right {
		for left < right && arr[right] >= arr[0] {
			right--
		}
		for left < right && arr[left] <= arr[0] {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[0] = arr[left]
	arr[left] = tmp
	quickSort2(arr[:left])
	quickSort2(arr[left+1:])
}

func main() {
	arr := []int{5, 8, 3, 2, 1}
	quickSort2(arr)
	fmt.Println(arr)
}
