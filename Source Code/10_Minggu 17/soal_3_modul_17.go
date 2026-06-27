package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())

	var N int
	fmt.Scan(&N)

	var a, b, c, d int

	for i := 0; i < N; i++{
		x := rand.Float64()
		y := rand.Float64()

		if x < 0.5{
			if y < 0.5{
				a++
			}else{
				c++
			}
		}else{
			if y < 0.5{
				b++
			}else{
				d++
			}
		}
	}

	curahA := float64(a) * 0.0001
	curahB := float64(b) * 0.0001
	curahC := float64(c) * 0.0001
	curahD := float64(d) * 0.0001

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahD)
}