package main

import "fmt"

func main() {
	array := []int{1,7,9,-1,-11,12,17,9}
	//array := []int{-1, -11}
	quickSort(&array, 0, len(array)-1)
}

func quickSort(array *[]int, start, end int) {
	if start < end {
		pivot := partition(array, start, end)
		fmt.Printf("pivot:%d, start:%d, end:%d, array:%v\n", pivot, start, end, *array)

		quickSort(array, start, pivot-1)
		quickSort(array, pivot+1, end)
	}
}

func partition(array *[]int, start, end int) int {
	if start >= end {
		return 0
	}

	pivot := start
	start++
	for start < end {
		for (*array)[end] >= (*array)[pivot] && start < end {
			end--
		}

		for (*array)[start] <= (*array)[pivot] && start < end {
			start++
		}

		(*array)[start], (*array)[end] = (*array)[end], (*array)[start]
	}

	if (*array)[start] < (*array)[pivot] {
		(*array)[start], (*array)[pivot] = (*array)[pivot], (*array)[start]
	}

	return start
}
