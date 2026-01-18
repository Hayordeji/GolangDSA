package main

import "fmt"

func main() {
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(0)) // 1
}
func factorial(n int) int {
	// Base case
	if n <= 1 {
		return 1
	}

	// Recursive case
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	// Base cases
	if n <= 1 {
		return n
	}

	// Recursive case
	return fibonacci(n-1) + fibonacci(n-2)
}

// WRONG IMPLEMENTATION
func countdown(n int) {
	fmt.Println(n)
	countdown(n - 1) //It never stops because of no present base case
}

// Infinite recursion - n never changes
func bad(n int) int {
	if n == 0 {
		return 0
	}
	return bad(n) // Problem doesn't get smaller
}
