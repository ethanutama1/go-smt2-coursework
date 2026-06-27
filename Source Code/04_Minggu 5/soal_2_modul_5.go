package main
import "fmt"

func segitigaSikuSiku(n int) {
	if n == 0 {
		return
	}
	segitigaSikuSiku(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func main(){
	// deklarasi segitiga siku siku bintang menggunakan rekursif
	segitigaSikuSiku(5)
}