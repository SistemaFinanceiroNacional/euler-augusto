package main

func Euler2() {
	_ = Fib(10)
}

func Fib(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 2
	}

	return Fib(n-1) + Fib(n-2)
}

// Escrever uma função que recebe um número e retorna a soma dos números fibonacci pares cujo o valor é menor que o número dado.
