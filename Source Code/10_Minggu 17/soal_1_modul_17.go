package main

import "fmt"

func main(){
	var data float64
	var jumlah float64
	var banyak int

	fmt.Scan(&data)

	if data == 9999{
		fmt.Println("Tidak ada data yang dimasukkan")
		return
	}

	jumlah = data
	banyak = 1

	fmt.Scan(&data)

	for data != 9999{
		jumlah += data
		banyak++
		fmt.Scan(&data)
	}

	rerata := jumlah / float64(banyak)
	fmt.Printf("Rata-rata = %.2f", rerata)
}