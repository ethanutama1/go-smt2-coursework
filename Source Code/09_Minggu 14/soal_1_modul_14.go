//Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap". 
package main

import (
	"fmt"
)
func main() {
	var data []int
	var input int
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}
	insertionSort(data)
	fmt.Println(data)
	if len(data) < 2 {
		fmt.Println("data berjarak tidak tetap")
		return
	}
	distance := data[1] - data[0]
	for i := 2; i < len(data); i++ {
		if data[i]-data[i-1] != distance {
			fmt.Println("data berjarak tidak tetap")
			return
		}
	}
	fmt.Printf("data berjarak %d\n", distance)
}
func insertionSort(arr []int){
		for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}