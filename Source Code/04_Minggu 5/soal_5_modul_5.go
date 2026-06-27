package main

import "fmt"

func pangkat(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * pangkat(base, exponent-1)
}

func main() {
	// membuat sistem pemangkatan dengan rekursif
	var base, exponent int
	fmt.Print("Masukkan bilangan dasar dan pangkatnya (misal: 2 3): ")
	fmt.Scan(&base, &exponent)
	fmt.Printf("%d pangkat %d adalah %d\n", base, exponent, pangkat(base, exponent))
}
