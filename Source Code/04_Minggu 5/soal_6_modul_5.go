package main
import "fmt"
func barisanGanjil(n int) {
	if n == 0 {
		return
	}
	barisanGanjil(n - 1)
	fmt.Printf("%d ", 2*(n-1)+1)
}
func main (){
	// membuat sistem barisan ganjil dengan rekursif
	var n int
	fmt.Print("Masukkan jumlah bilangan ganjil yang ingin ditampilkan: ")
	fmt.Scan(&n)
	fmt.Printf("Barisan ganjil hingga %d: ", n)
	barisanGanjil(n)
	fmt.Println()
}