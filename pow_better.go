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

func myPow(x float64, n int) float64{
	if n == 0 {
		return 1
	}

	if x == 1 || n == 1{
		return x
	}

	if n < 0 {
		x = 1/x
		n = -n
	}

	pow := float64(1)
	for n != 0 {
		if n&1 == 1 {
			pow *= x
		}

		x *= x
		n>>=1 
	}

	return pow
}