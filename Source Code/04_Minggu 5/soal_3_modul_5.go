package main
import "fmt"
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func main(){
	// membuat sistem pemfaktoran bilangan dengan rekursif
	var n int
	fmt.Print("Masukkan bilangan untuk dihitung faktorialnya: ")
	fmt.Scan(&n)
	result := faktorial(n)
	fmt.Printf("Faktorial dari %d adalah %d\n", n, result)
}
