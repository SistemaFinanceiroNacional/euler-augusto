package main

import "fmt"

func Euler2() {
	SumEvenValuedFibonacciTerms(4000000)
}

func FibRecursion(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 2
	}

	return FibRecursion(n-1) + FibRecursion(n-2)
}

func SumEvenValuedFibonacciTerms(n int) {
	var sumResult, index int

	for {
		fibResult := FibRecursion(index)
		index++
		if fibResult%2 == 0 {
			sumResult += fibResult
		}
		if fibResult >= n {
			break
		}
	}
	fmt.Println("Euler 2 =", sumResult)
}
