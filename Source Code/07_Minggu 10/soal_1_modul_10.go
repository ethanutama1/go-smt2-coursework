package main
import "fmt"
type Kelinci struct {
	berat float64
}
func main() {
	var N int
	fmt.Scan(&N)
	var kelinci [1000]Kelinci
	for i := 0; i < N; i++ {
		fmt.Scan(&kelinci[i].berat)
	}
	var min, max float64
	min = kelinci[0].berat
	max = kelinci[0].berat
	for i := 1; i < N; i++ {
		if kelinci[i].berat < min {
			min = kelinci[i].berat
		}
		if kelinci[i].berat > max {
			max = kelinci[i].berat
		}
	}
	fmt.Printf("%.2f %.2f\n", min, max)
}