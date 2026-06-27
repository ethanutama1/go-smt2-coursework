package domino

import (
	"math/rand"
	"time"
)

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
