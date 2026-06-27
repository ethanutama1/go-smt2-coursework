package main
import "fmt"
func BarisanBilangan(n int) int {
	if n == 1 {
		return 1
	}
	return n + BarisanBilangan(n-1)
}
func main(){
	// mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.
	// contoh input : 5 output : 5 4 3 2 1 2 3 4 5  
	var n int 
	fmt.Print("Masukkan jumlah bilangan dalam barisan: ")
	fmt.Scan(&n)
	fmt.Printf("Barisan bilangan hingga %d: ", n)
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	for i := 2; i <= n; i++ {
		fmt.Print(i," ")
	}
	fmt.Println()
}