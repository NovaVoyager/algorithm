package main

import "fmt"

func main() {
	fmt.Println(fibonacci2(3))
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
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

	n1, n2 := 0, 1

	for i := 2; i <= n; i++ {

		n1, n2 = n2, n1+n2

	}

	return n2

}
