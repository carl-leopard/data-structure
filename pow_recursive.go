package main

func main() {
	return myPow(1, 100)
}

//n is not negtive number
func myPow(x int, n int) int {
	if n == 0 {
		if x >= 0 {
			return 1
		}

		return -1
	}

	if n == 1 {
		return x
	}

	y := myPow(x, n/2)

	if n&1 == 0 {
		return y * y
	}

	return y * y * x
}
