package main

import "fmt"

func main() {
	fmt.Println(fibonacci2(5))
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci2(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 1, 1

	for i := 3; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}
