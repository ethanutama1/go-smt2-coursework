package main

import "fmt"

func main() {
	var K int
	fmt.Print("Nilai K: ")
	fmt.Scan(&K)

	var hasil float64

	for k := 0; k <= K; k++ {
		hasil = ((4*float64(k) + 2) * (4*float64(k) + 2)) /
			((4*float64(k) + 1) * (4*float64(k) + 3))
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}