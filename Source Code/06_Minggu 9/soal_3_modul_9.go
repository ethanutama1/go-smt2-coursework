package main

import (
	"fmt"
)
func main() {
	var club1, club2 string
	var score1, score2 int
	var winners []string
	fmt.Print("Masukkan nama klub pertama: ")
	fmt.Scan(&club1)
	fmt.Print("Masukkan nama klub kedua: ")
	fmt.Scan(&club2)
	for {
		fmt.Print("Masukkan skor klub pertama: ")
		fmt.Scan(&score1)
		fmt.Print("Masukkan skor klub kedua: ")
		fmt.Scan(&score2)
		if score1 < 0 || score2 < 0 {
			fmt.Println("Skor tidak valid. Proses input dihentikan.")
			break
		}
		if score1 > score2 {
			winners = append(winners, club1)
		} else if score2 > score1 {
			winners = append(winners, club2)
		}
	}
	fmt.Println("Daftar klub yang memenangkan pertandingan:")
	for _, winner := range winners {
		fmt.Println(winner)
	}	
}