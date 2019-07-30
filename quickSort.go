package main

import (
	"fmt"
)

func main() {
	array := []int{1, 11, 7, 5, 12, 3, -2, 4}
	quickSort(&array, 0, len(array)-1)
	fmt.Println(array)
}

func quickSort(array *[]int, start, end int) {
	if start < end {
		pivot := partitionSort(array, start, end)
		fmt.Printf("pivot:%d, start:%d, end:%d, array:%v\n", pivot, start, end, *array)
		quickSort(array, start, pivot-1)
		quickSort(array, pivot+1, end)
	}
}

func partitionSort(array *[]int, start, end int) int {
	if end <= start {
		return 0
	}

	pivot := (*array)[start]
	head := start + 1
	tail := end

	for head < tail {
		for head < tail && (*array)[tail] >= pivot {
			tail--
		}

		for head < tail && (*array)[head] <= pivot {
			head++
		}

		(*array)[head], (*array)[tail] = (*array)[tail], (*array)[head]
	}

	if (*array)[head] > pivot {
		(*array)[head], (*array)[start] = (*array)[start], (*array)[head]
	}

	return head
}
