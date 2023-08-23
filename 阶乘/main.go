package main

import "fmt"

func main() {
	res := factorial(2)
	fmt.Println(res)
}

func factorial(n int) int {
	sum := 1

	for i := 1; i <= n; i++ {
		sum *= i
	}

	return sum
}

func factorialRecur(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	// 从 1 个分裂出 n 个
	for i := 0; i < n; i++ {
		count += factorialRecur(n - 1)
	}
	return count
}
