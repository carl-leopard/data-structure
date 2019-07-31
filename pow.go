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

	if x == 1 {
		return float64(x)
	}

	multi := x
	i := 2
	negative := false

	if n < 0 {
		negative = true
		n = -n
	}
	for i <= n{
		multi *= multi
		i *= 2
	}

	if n&1 != 0 && n != 1{
		multi *= x
	}

	if negative {
		return float64(1)/float64(multi)
	}

	return float64(multi)
}