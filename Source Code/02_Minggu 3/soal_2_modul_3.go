package main

import "fmt"

func F(x int) int {
	hasil := x * x
	return hasil
}

func G(x int) int {
	hasil := x - 2
	return hasil
}

func H(x int) int {
	hasil := x + 1
	return hasil
}

func main() {
	var a, b, c int
	fmt.Print("Masukan 3 Bilangan bulat: ")
	fmt.Scan(&a, &b, &c)

	hasil1 := F(G(H(a)))

	hasil2 := G(H(F(b)))

	hasil3 := H(F(G(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}