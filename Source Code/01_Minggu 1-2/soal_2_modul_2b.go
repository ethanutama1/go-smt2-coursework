package main

import "fmt"

func main() {
	var a, b float64

	for {
		fmt.Print("Masukan berat belanjaan: ")
		fmt.Scan(&a, &b)

		if a+b > 150 || a < 0 || b < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		oleng := (a-b >= 9) || (b-a >= 9)
		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}
}