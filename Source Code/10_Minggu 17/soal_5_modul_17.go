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

	inCircle := 0

	for i := 0; i < N; i++{
		x := rand.Float64()
		y := rand.Float64()
		
		dx := x - 0.5
		dy := y - 0.5

		if dx*dx+dy*dy <= 0.25{
			inCircle++
		}
	}

	piEstimate := 4.0 * float64(inCircle) / float64(N)
	fmt.Printf("Banyak Topping: %d Topping pada Pizza: %d PI : %.10f\n", N, inCircle, piEstimate)
}