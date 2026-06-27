package main

import (
	"fmt"
	"math"
)

func main(){

	N := 1000000
	sum := 0.0 
	for i := 1; i <= N; i++{
		sign := 1.0 
		if i%2 == 0{
			sign  = -1.0
		}
		term := sign / float64(2*i-1)
		sum += term
	}

	pil1 := 4 * sum
	fmt.Printf("N suku pertama: %d Hasil PI: %.7f\n", N, pil1)

	sumPrev := 0.0
	i := 1
	for {
		sign := 1.0
		if i%2 == 0{
			sign = -1.0
		}

		term := sign / float64(2*i-1)
		sumCurrent := sumPrev + term

		if i > 1{
			diff := math.Abs(sumCurrent - sumPrev)
			if diff <= 0.00001{
				piPrev := 4 * sumPrev
				piCurrent := 4 * sumCurrent
				fmt.Printf("Hasil PI: %.10f\n", piPrev)
				fmt.Printf("Hasil PI: %.10f\n", piCurrent)
				fmt.Printf("Pada i ke: %d\n", i)
				break
			}
		}

		sumPrev = sumCurrent
		i++
	}
}