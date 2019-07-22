package main

//https://leetcode-cn.com/problems/sliding-window-maximum

//windows[0] 永远是最大的

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0, 0)
	}

    result := make([]int, 0, len(nums)-k+1)
    windows := make([]int, 0, k+1)


	kGroupMaxIndex := k-1
	for i, n :=range nums {
		
		windows = append(windows, n)
		//如果新增元素 >= windows[0]，保持 windows[0] 永远是最大的
		if n >= windows[0] {
			windows = windows[len(windows)-1:]
		}

		//新增元素之后， 窗口元素数量 > K
		if len(windows) > k {
			windows = windows[1:]

			//保持 windows[0] 永远是最大的
			//moveMaxToFront(&windows)
			if len(windows) > 1 {
				maxIndex := 0
				i := 1 
				for i < len(windows) {
					if windows[i] >= windows[maxIndex] {
						maxIndex = i
					}
			        
			        i++
				}

			    windows = windows[maxIndex:]
			}
		}

		if i >= kGroupMaxIndex {
			result = append(result, windows[0])
		}
	}

	return result
}

func moveMaxToFront(windows *[]int) {
    if len(*windows) <= 1 {
        return
    }
    
	maxIndex := 0
	i := 1 
	for i < len(*windows) {
		if (*windows)[i] >= (*windows)[maxIndex] {
			maxIndex = i
		}
        
        i++
	}

    *windows = (*windows)[maxIndex:]
}