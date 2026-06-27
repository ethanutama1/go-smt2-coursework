//Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom. 
package main

import (
	"fmt"
)
const NMAX = 127 
type table [NMAX]rune
tab := table{}
m  := int
func isiArray(t *table, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
		if t[i] == '.' {
			break
		}
	}
}

func cetakArray(t *table, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
}

func balikanArray(t *table, n int)	{
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%c ", t[i])
	}
}

func main (){
	var tab table
	var n int

	fmt.Print("Masukkan jumlah karakter (maksimal 127): ")
	fmt.Scan(&n)

	isiArray(&tab, n)
	balikanArray(&tab, n)
	cetakArray(&tab, n)
}