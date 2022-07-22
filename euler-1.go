package main

import "fmt"

func Euler1() {

	n1 := 3
	n2 := 5
	var result int

	for i := 1; i < 1000; i++ {
		if i%n1 == 0 || i%n2 == 0 {
			result += i
		}
	}

	fmt.Println("Euler 1 =", result)
}
