package main

import "fmt"

//转载 https://zhuanlan.zhihu.com/p/45725214
//数组下标从 0 开始

func main() {
	array := []int{57, 40, 38, 11, 13, 34, 48, 75, 6, 19, 9, 7}
	fmt.Println("origin: ", array)

	buildHeap(&array)
	fmt.Println("after build: ", array)

	for i := len(array)-1; i>=0; i-- {
		sortHeap(&array, i)
		fmt.Println("exchange ", i, ": ", array)
	}

	fmt.Println("sorted: ", array)
}

//max heap
func buildHeap(array *[]int) {
	var (
		left int
		right int
	)
	for parent := len(*array) /2 - 1 ; parent >= 0; parent-- {
		left = 2 * parent + 1
		right = left + 1
		if left <= len(*array)-1 && right <= len(*array)-1 {
			if (*array)[left] > (*array)[parent] && (*array)[left] > (*array)[right] {
				(*array)[left], (*array)[parent] = (*array)[parent], (*array)[left]
				continue
			}

			if (*array)[right] > (*array)[parent] && (*array)[right] > (*array)[left] {
				(*array)[right], (*array)[parent] = (*array)[parent], (*array)[right]
				continue
			}

			continue
		}

		if left <= len(*array)-1 {
			if (*array)[left] > (*array)[parent]  {
				(*array)[left], (*array)[parent] = (*array)[parent], (*array)[left]
			}
		}
	}
}

func sortHeap(array *[]int, index int) {
	(*array)[index], (*array)[0] = (*array)[0], (*array)[index]
	adjustHeap(array, index-1)
}

func adjustHeap(array *[]int, maxIndex int) {
	var (
		left int
		right int
	)

	for parent := 0; parent <= maxIndex; {
		left = 2 * parent + 1
		right = left + 1
		if left <= maxIndex && right <= maxIndex {
			if (*array)[left] > (*array)[parent] && (*array)[left] > (*array)[right] {
				(*array)[left], (*array)[parent] = (*array)[parent], (*array)[left]

				parent = left
				continue
			}

			if (*array)[right] > (*array)[parent] && (*array)[right] > (*array)[left] {
				(*array)[right], (*array)[parent] = (*array)[parent], (*array)[right]

				parent = right
				continue
			}

			break
		}

		if left <= maxIndex {
			if (*array)[left] > (*array)[parent]  {
				(*array)[left], (*array)[parent] = (*array)[parent], (*array)[left]

				break
			}
		}

		break
	}
}

