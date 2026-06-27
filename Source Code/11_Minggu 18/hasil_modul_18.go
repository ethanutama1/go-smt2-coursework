package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ========== BAGIAN DOMINO ==========

type Domino struct {
	Sisi1 int
	Sisi2 int
	Balak bool
}

type Dominoes struct {
	Kartu [28]Domino
	Sisa  int
}

func InisialisasiDomino() Dominoes {
	var d Dominoes
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d.Kartu[idx] = Domino{Sisi1: i, Sisi2: j, Balak: i == j}
			idx++
		}
	}
	d.Sisa = 28
	return d
}

func KocokKartu(d *Dominoes) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(d.Sisa, func(i, j int) {
		d.Kartu[i], d.Kartu[j] = d.Kartu[j], d.Kartu[i]
	})
}

func AmbilKartu(d *Dominoes) Domino {
	if d.Sisa == 0 {
		return Domino{-1, -1, false}
	}
	kartu := d.Kartu[d.Sisa-1]
	d.Sisa--
	return kartu
}

func GambarKartu(d Domino, suit int) int {
	if d.Sisi1 == suit {
		return d.Sisi2
	} else if d.Sisi2 == suit {
		return d.Sisi1
	}
	return -1
}

func NilaiKartu(d Domino) int {
	return d.Sisi1 + d.Sisi2
}

func GaliKartu(d *Dominoes, target Domino) Domino {
	for d.Sisa > 0 {
		k := AmbilKartu(d)
		if k.Sisi1 == target.Sisi1 || k.Sisi1 == target.Sisi2 ||
			k.Sisi2 == target.Sisi1 || k.Sisi2 == target.Sisi2 {
			return k
		}
	}
	return Domino{-1, -1, false}
}

func SepasangKartu(d1 Domino, d2 Domino) bool {
	return (NilaiKartu(d1) + NilaiKartu(d2)) == 12
}

// ========== BAGIAN KARAKTER ==========

var pita string
var indeks int
var currentChar byte

func Start(input string) {
	pita = input
	indeks = 0
	if len(pita) > 0 {
		currentChar = pita[indeks]
	}
}

func Maju() {
	if !Eop() {
		indeks++
		currentChar = pita[indeks]
	}
}

func Eop() bool {
	return currentChar == '.'
}

func Cc() byte {
	return currentChar
}

func ProsesMesinKarakter(input string) {
	Start(input)

	totalKarakter := 0
	jumlahA := 0
	jumlahLE := 0
	var prevChar byte = ' '

	for !Eop() {
		c := Cc()
		totalKarakter++

		if c == 'A' {
			jumlahA++
		}

		if prevChar == 'L' && c == 'E' {
			jumlahLE++
		}

		prevChar = c
		Maju()
	}

	frekuensiA := 0.0
	if totalKarakter > 0 {
		frekuensiA = float64(jumlahA) / float64(totalKarakter)
	}

	fmt.Printf("Total karakter terbaca: %d\n", totalKarakter)
	fmt.Printf("Jumlah huruf 'A': %d\n", jumlahA)
	fmt.Printf("Frekuensi 'A' terhadap seluruh karakter: %.2f\n", frekuensiA)
	fmt.Printf("Jumlah kemunculan kata 'LE': %d\n", jumlahLE)
}

// ========== MAIN ==========

func main() {
	fmt.Println("===================================")
	fmt.Println("      PENGUJIAN MESIN DOMINO       ")
	fmt.Println("===================================")

	tumpukan := InisialisasiDomino()
	fmt.Printf("Jumlah kartu awal: %d\n", tumpukan.Sisa)

	KocokKartu(&tumpukan)
	fmt.Println("Kartu telah dikocok!")

	kartuKu := AmbilKartu(&tumpukan)
	fmt.Printf("Kartu yang diambil: (%d, %d), Balak: %t\n", kartuKu.Sisi1, kartuKu.Sisi2, kartuKu.Balak)
	fmt.Printf("Sisa kartu di tumpukan: %d\n", tumpukan.Sisa)

	fmt.Printf("Total nilai kartuKu: %d\n", NilaiKartu(kartuKu))

	suitTarget := 3
	hasilSuit := GambarKartu(kartuKu, suitTarget)
	if hasilSuit != -1 {
		fmt.Printf("Kartu cocok dengan suit %d, sisi lainnya adalah %d\n", suitTarget, hasilSuit)
	} else {
		fmt.Printf("Kartu tidak memiliki suit %d\n", suitTarget)
	}

	fmt.Println("\n--- Menggali Kartu ---")
	kartuCocok := GaliKartu(&tumpukan, kartuKu)
	if kartuCocok.Sisi1 != -1 {
		fmt.Printf("Berhasil menggali kartu yang cocok: (%d, %d)\n", kartuCocok.Sisi1, kartuCocok.Sisi2)
	} else {
		fmt.Println("Tumpukan habis, tidak ada kartu yang cocok.")
	}

	fmt.Println("\n--- Mengecek Sepasang Kartu ---")
	kartuA := Domino{Sisi1: 6, Sisi2: 6, Balak: true}
	kartuB := Domino{Sisi1: 0, Sisi2: 0, Balak: true}
	fmt.Printf("Apakah kartu (6,6) dan (0,0) bernilai total 12? %t\n", SepasangKartu(kartuA, kartuB))

	fmt.Println("\n\n===================================")
	fmt.Println("     PENGUJIAN MESIN KARAKTER      ")
	fmt.Println("===================================")

	inputPita := "PASRAHKAN SAJA DAN BERDOA"
	fmt.Printf("Pita karakter yang dibaca: \"%s\"\n\n", inputPita)
	fmt.Println("Hasil Analisis:")

	ProsesMesinKarakter(inputPita)
}