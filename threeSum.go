package main

import "fmt"

func main() {
	array := []int{1,2,6,5,3,4}
	//array := []int{-1, -11}
	fmt.Println(threeSum(&array, 15))


}

func threeSum(array *[]int, sum int) []int{
	quickSort(array, 0, len(*array)-1)

	result := make([]int, 0, 3)


	for i := 0; i < len(*array)-2; i++ {
		start := 0
		end := len(*array) -1

		result = result[:0]
		result = append(result, i)
		twoSum := sum - (*array)[i]

		//for start < end {
			for (start == i || (*array)[start] + (*array)[end] < twoSum) && start < end {
				start++
			}

			fmt.Printf("start. i:%d, start:%d, end:%d\n", i, start, end)
			if start == end {
				continue
			}

			for (end == i || (*array)[start] + (*array)[end] > twoSum) && start < end{
				end--
			}
			if start == end {
				continue
			}

			fmt.Printf("end. i:%d, start:%d, end:%d\n", i, start, end)

			if (*array)[start] + (*array)[end] == twoSum {
				result = append(result, start, end)
				return result
			}
		//}
	}
	return result
}

func quickSort(array *[]int, start, end int) {
	if start < end {
		pivot := partition(array, start, end)
		fmt.Printf("pivot:%d, start:%d, end:%d, array:%v\n", pivot, start, end, *array)

		//对中轴线的左部分进行排序
		quickSort(array, start, pivot-1)

		//对中轴线的右部分进行排序
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
		//必须从最右边开始
		for (*array)[end] >= (*array)[pivot] && start < end {
			end--
		}

		for (*array)[start] <= (*array)[pivot] && start < end {
			start++
		}

		//进行交换
		(*array)[start], (*array)[end] = (*array)[end], (*array)[start]
	}

	//循环退出时 判断是否需要交换
	if (*array)[start] < (*array)[pivot] {
		(*array)[start], (*array)[pivot] = (*array)[pivot], (*array)[start]
	}

	return start
}