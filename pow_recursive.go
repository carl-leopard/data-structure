package main

import "fmt"

func main() {
	fmt.Println("myPow(1, 1):", myPow(1, 1))
	fmt.Println("myPow(1, 0):", myPow(1, 0))
	fmt.Println("myPow(1, -1):", myPow(1, -1))

	fmt.Println("myPow(1, 0):", myPow(1, 0))
	fmt.Println("myPow(-1, 0):",myPow(-1, 0))

	fmt.Println("myPow(2, 4):", myPow(2, 4))
	fmt.Println("myPow(2, 5):",myPow(2, 5))
	fmt.Println("myPow(2, 0):", myPow(2, 0))
	fmt.Println("myPow(2, 1):", myPow(2, 1))
	fmt.Println("myPow(2, -1):", myPow(2, -1))
	fmt.Println("myPow(2, -2):", myPow(2, -2))
	fmt.Println("myPow(2, -3):", myPow(2, -3))
}

func myPow(x int, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return float64(x)
	}

	if n < 0 {
		return 1/myPow(x, -n)
	}

	y := myPow(x, n/2)

	if n&1 == 0 {
		return float64(y) * float64(y)
	}

	return float64(y) * float64(y) * float64(x)
}
