package karakter

import "fmt"

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
