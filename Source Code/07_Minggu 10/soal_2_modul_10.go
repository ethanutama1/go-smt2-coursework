package main
import "fmt"
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	var ikan [1000]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}
	var wadah [1000]float64
	var totalWadah int
	for i := 0; i < x; i++ {
		wadah[totalWadah] += ikan[i]
		if (i+1)%y == 0 {
			totalWadah++
		}
	}
	for i := 0; i < totalWadah; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
	fmt.Println()
	var rataRata float64
	for i := 0; i < totalWadah; i++ {
		rataRata += wadah[i]
	}	
	rataRata /= float64(totalWadah)
	fmt.Printf("%.2f\n", rataRata)
}