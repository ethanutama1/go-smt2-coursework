package main
import "fmt"
func VolumeKubus(a int) int {
	return a * a * a
}
func LuasKubus(a int) int {
	return 6 * a * a
}
func VolumeBalok(p int, l int, t int) int {
	return p * l * t
}
func LuasBalok(p int, l int, t int) int {
	return 2 * (p * l + p * t + l * t)
}
func VolumeTabung(r int, t int) float64 {
	return 3.14 * float64(r) * float64(r) * float64(t)
}
func LuasTabung(r int, t int) float64 {
	return 2 * 3.14 * float64(r) * (float64(r) + float64(t))
}
func main() {
	var a, p, l, t int

	fmt.Print("Masukan Sisi Kubus: ")
	fmt.Scan(&a)
	fmt.Print("Masukan Panjang Balok: ")
	fmt.Scan(&p)
	fmt.Print("Masukan Lebar Balok: ")
	fmt.Scan(&l)
	fmt.Print("Masukan Tinggi Balok: ")
	fmt.Scan(&t)
	fmt.Print("Masukan Jari-jari Tabung: ")
	fmt.Scan(&a)
	fmt.Print("Masukan Tinggi Tabung: ")
	fmt.Scan(&t)

	volumeKubus := VolumeKubus(a)
	luasKubus := LuasKubus(a)
	volumeBalok := VolumeBalok(p, l, t)
	luasBalok := LuasBalok(p, l, t)
	volumeTabung := VolumeTabung(a, t)
	luasTabung := LuasTabung(a, t)

	fmt.Printf("Volume Kubus dengan sisi %d adalah %d\n", a, volumeKubus)
	fmt.Printf("Luas Kubus dengan sisi %d adalah %d\n", a, luasKubus)
	fmt.Printf("Volume Balok dengan panjang %d, lebar %d, dan tinggi %d adalah %d\n", p, l, t, volumeBalok)
	fmt.Printf("Luas Balok dengan panjang %d, lebar %d, dan tinggi %d adalah %d\n", p, l, t, luasBalok)
	fmt.Printf("Volume Tabung dengan jari-jari %d dan tinggi %d adalah %.2f\n", a, t, volumeTabung)
	fmt.Printf("Luas Tabung dengan jari-jari %d dan tinggi %d adalah %.2f\n", a, t, luasTabung)


}