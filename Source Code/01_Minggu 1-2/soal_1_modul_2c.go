package main

import "fmt"

func main() {
	var gram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	kg := gram / 1000
	sisa := gram % 1000

	biaya := kg * 10000

	if kg <= 10 {
		if sisa >= 500 {
			biaya += sisa * 5
		} else {
			biaya += sisa * 15
		}
	}

	fmt.Println("Total biaya:", biaya)
}