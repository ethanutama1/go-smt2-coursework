package main
import "fmt"
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main(){
	// membuat sistem deret fibonacci dengan rekursif
	var n int
	fmt.Print("Masukkan jumlah deret Fibonacci yang ingin ditampilkan: ")
	fmt.Scan(&n)
	fmt.Printf("Deret Fibonacci hingga %d: ", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}