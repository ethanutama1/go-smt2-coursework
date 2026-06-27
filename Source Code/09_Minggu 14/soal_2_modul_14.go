// Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini: 
//Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari. 
package main
import (
	"fmt"
)
type Buku struct {
	id int
	Judul string
	Penulis string
	penerbit string
	tahunTerbit int
	eksemplar int
	rating int
}
type DaftarBuku struct {
	buku []Buku
}
var pustaka = DaftarBuku{
	buku: []Buku{
		{id: 1, Judul: "Buku A", Penulis: "Penulis A", penerbit: "Penerbit A", tahunTerbit: 2020, eksemplar: 5, rating: 4},
		{id: 2, Judul: "Buku B", Penulis: "Penulis B", penerbit: "Penerbit B", tahunTerbit: 2019, eksemplar: 3, rating: 5},
			{id: 3, Judul: "Buku C", Penulis: "Penulis C", penerbit: "Penerbit C", tahunTerbit: 2021, eksemplar: 2, rating: 3},
		},
		
	}

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var id int
		var Judul, Penulis, penerbit string
		var tahunTerbit, eksemplar, rating int
		fmt.Scan(&id, &Judul, &Penulis, &penerbit, &tahunTerbit, &eksemplar, &rating)
		buku := Buku{id: id, Judul: Judul, Penulis: Penulis, penerbit: penerbit, tahunTerbit: tahunTerbit, eksemplar: eksemplar, rating: rating}
		pustaka.buku = append(pustaka.buku, buku)
	}
	var ratingCari int
	fmt.Scan(&ratingCari)
	for _, buku := range pustaka.buku {
		if buku.rating == ratingCari {
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun Terbit: %d, Eksemplar: %d, Rating: %d\n", buku.id, buku.Judul, buku.Penulis, buku.penerbit, buku.tahunTerbit, buku.eksemplar, buku.rating)
		}
	}
}