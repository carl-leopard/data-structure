package main

import "fmt"

func main() {
	array := []int{2, 4, 10, 23, 7}
	fmt.Println(array)
	array = insertByIndex(array, 0, 9)
	fmt.Println(array)
	array = insertByIndex(array, len(array), 9)
	fmt.Println(array)
	array = insertByIndex(array, 5, 9)
	fmt.Println(array)
	fmt.Println()

	array = []int{2, 4, 10}
	fmt.Println(array)
	fmt.Println("len(keys) = 0: ", getByRange([]int{}, 0, len(array)))
	fmt.Println("start > end: ", getByRange(array, 2, 1))
	fmt.Println("start < 0: ", getByRange(array, -1, len(array)))
	fmt.Println("start >= len(keys): ", getByRange(array, len(array), len(array)-1))
	fmt.Println("end >= len(keys): ", getByRange(array, 1, len(array)))
	fmt.Println()
	fmt.Println("start = 0: ", getByRange(array, 0, len(array)-1))
	fmt.Println("start = end: ", getByRange(array, 0, 0))
	fmt.Println("start = end: ", getByRange(array, len(array)-1, len(array)-1))
}

func insertByIndex(keys []int, index int, val int) []int {
	if index >= len(keys) {
		keys = append(keys, val)
		return keys
	}

	//keys = append([]int{val}, keys...)
	if index == 0 {
		arr := make([]int, 0, len(keys)+1)
		arr = append(arr, val)
		arr = append(arr, keys...)
		return arr
	}

	//keys = append(keys[:index], append([]int{val}, keys[index:]))
	arr := make([]int, len(keys)-index)
	copy(arr, keys[index:])
	keys = append(keys[:index], val)
	keys = append(keys, arr...)
	return keys
}

//左闭右开
func getByRange(keys []int, start, end int) []int {
	if len(keys) == 0 ||
		start > end ||
		start < 0 ||
		start >= len(keys) ||
		end >= len(keys) {
		return make([]int, 0, 0)
	}

	dst := make([]int, end-start, end-start+1)
	copy(dst, keys[start:end])
	dst = append(dst, keys[end])
	return dst
}
