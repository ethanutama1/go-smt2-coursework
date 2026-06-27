package main

import "fmt"

func main(){
	var x string
	var n int
	var data string

	fmt.Scan(&x)
	fmt.Scan(&n)

	posisiPertama := 0 
	jumlah := 0

	for i := 1; i <= n; i++{
		fmt.Scan(&data)
		if data == x {
			jumlah++
			if posisiPertama == 0{
				posisiPertama = i
			}
		}
	}

	if jumlah > 0{
		fmt.Println("String x ditemukan")
		fmt.Printf("Posisi pertama: %d\n", posisiPertama)
		fmt.Printf("Jumlah kemunculan: %d\n", jumlah)
		if jumlah >= 2{
			fmt.Println("Terdapat setidaknya dua string x")
		}else{
			fmt.Println("Hanya terdapat satu string x")
		}
	}else{
		fmt.Println("String x tidak ditemukan")
	}
}




