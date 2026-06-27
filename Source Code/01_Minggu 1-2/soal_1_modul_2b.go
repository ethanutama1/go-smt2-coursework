package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	berhasil := true
	i := 1

	for i <= 5 {
		fmt.Scan(&w1, &w2, &w3, &w4)

		if !(w1 == "merah" && w2 == "kuning" && 
		    w3 == "hijau" && w4 == "ungu") {
			berhasil = false
		}
		i++
	}

	fmt.Println("BERHASIL:", berhasil)
}