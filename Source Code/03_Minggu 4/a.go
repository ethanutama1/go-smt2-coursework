package main 
import "fmt"
func factorial(n int)int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
func main(){
	var x, hasil int
	fmt.Print("Masukan angka yang akan di factorial: ")
	fmt.Scan(&x)
	hasil = factorial(x)
	fmt.Print("Hasil factorial dari ", x, " adalah : ", hasil)
}