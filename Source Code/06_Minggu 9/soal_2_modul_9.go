package main

import (
	"fmt"
	"math"
)
func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)
	array := make([]int, N)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("a. Isi keseluruhan array:")
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println("\nb. Elemen dengan indeks ganjil:")
	for i := 1; i < N; i += 2 {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println("\nc. Elemen dengan indeks genap:")
	for i := 0; i < N; i += 2 {
		fmt.Printf("%d ", array[i])
	}
	var x int
	fmt.Print("\nd. Masukkan bilangan x untuk menampilkan elemen dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", array[i])
		}
	}
	var deleteIndex int
	fmt.Print("\ne. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&deleteIndex)
	if deleteIndex >= 0 && deleteIndex < N {
		array = append(array[:deleteIndex], array[deleteIndex+1:]...)
		N--
		fmt.Println("Isi array setelah penghapusan:")
		for i := 0; i < N; i++ {
			fmt.Printf("%d ", array[i])
		}
	} else {
		fmt.Println("Indeks tidak valid.")
	}
	sum := 0
	for i := 0; i < N; i++ {
		sum += array[i]
	}
	average := float64(sum) / float64(N)
	fmt.Printf("\nf. Rata-rata: %.2f\n", average)
	varianceSum := 0.0
	for i := 0; i < N; i++ {
		varianceSum += math.Pow(float64(array[i])-average, 2)
	}
	variance := varianceSum / float64(N)
	stdDev := math.Sqrt(variance)
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)
	var frequency int
	fmt.Print("h. Masukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&frequency)
	count := 0
	for i := 0; i < N; i++ {
		if array[i] == frequency {
			count++
		}
	}
	fmt.Printf("Frekuensi dari %d: %d\n", frequency, count)
}