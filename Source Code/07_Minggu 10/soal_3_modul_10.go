package main
import "fmt"
type arrBalita struct {
	berat float64
}
func rerata(arr []arrBalita, n int) float64 {	
	var total float64
	for i := 0; i < n; i++ {
		total += arr[i].berat
	}
	return total / float64(n)
}
func main() {
	var n int
	fmt.Scan(&n)
	var balita [1000]arrBalita
	for i := 0; i < n; i++ {
		fmt.Scan(&balita[i].berat)
	}
	var min, max float64
	min = balita[0].berat
	max = balita[0].berat
	for i := 1; i < n; i++ {
		if balita[i].berat < min {
			min = balita[i].berat
		}
		if balita[i].berat > max {
			max = balita[i].berat
		}
	}
	fmt.Printf("%.2f %.2f %.2f\n", min, max, rerata(balita[:n], n))
}