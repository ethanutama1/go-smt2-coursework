package main

import "fmt"

func factorial(n int) int {
	hasilFactorial := 1
	for i := n; i > 0; i-- {
		hasilFactorial *= i
	}
	return hasilFactorial

}
func permutation(n, r int) int {
	for i := n; i > n-r; i-- {
		n *= i
	}
	return n

}
func combination(n, r int) int {
	for i := n; i > n-r; i-- {
		n *= i
	}
	for i := r; i > 0; i-- {
		r *= i
	}
	return n / r
}
func main() {
	var n, r int
	fmt.Print("Enter n: ")
	fmt.Scanf("%d", &n)
	fmt.Print("Enter r: ")
	fmt.Scanf("%d", &r)
	fmt.Printf("Factorial of %d is %d\n", n, factorial(n, 1))
	fmt.Printf("Permutation of %d and %d is %d\n", n, r, permutation(n, r))
	fmt.Printf("Combination of %d and %d is %d\n", n, r, combination(n, r))
}
