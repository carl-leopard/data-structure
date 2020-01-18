package main

import "fmt"

func main() {
	array := []int{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
		//55, 94, 87, 1, 4,
		//, 55, 4,
	}
	fmt.Println(array)
	HeapSort(array)
	fmt.Println(array)
}

func HeapSort(array []int) {
	s := len(array)
	lastLeaf := s/2 - 1
	for i := lastLeaf; i >= 0; i-- {
		buildHeap(array, i, s-1)
	}
	fmt.Println(array)

	for i := s - 1; i >= 0; i-- {
		heapify(array, i)
	}
}

//使用特定区间的元素进行建堆（大顶堆）
//from i to end
//recursive, 调整后可能引起子节点不是大顶堆， 需要继续调整子节点
func buildHeap(array []int, i, end int) {
	left := i*2 + 1
	if left > end {
		return
	}

	maxIndex := findMaxIndex(array, i, end)
	if maxIndex == i {
		return
	}

	array[i], array[maxIndex] = array[maxIndex], array[i]
	buildHeap(array, maxIndex, end)
}

//将对顶元素与最后一个元素进行交换， 然后重新构建大顶堆
func heapify(array []int, end int) {
	array[0], array[end] = array[end], array[0]
	buildHeap(array, 0, end-1)
}

func findMaxIndex(array []int, i, end int) int {
	left := i*2 + 1
	right := left + 1

	max := i
	if left <= end && array[left] > array[max] {
		max = left
	}

	if right <= end && array[right] > array[max] {
		max = right
	}

	return max
}
