package main

import "fmt"
func main (){
	var r float64
	const pi = 3.1415926535

	fmt.Print("Masukan Jari-Jari: ")
	fmt.Scan(&r)

	Volume := (4/3)*pi*(r*r*r)
	Luas_Bola := 4*pi*(r*r)
	fmt.Println("Volumenya adalah: ", Volume)
	fmt.Println("Luasnya adalah: ", Luas_Bola)
}