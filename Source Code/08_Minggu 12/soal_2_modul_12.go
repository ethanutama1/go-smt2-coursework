// Berdasarkan program sebelumnya, buat program pilkart yang mencari siapa pemenang pemilihan ketua RT. Sekaligus juga ditentukan bahwa wakil ketua RT adalah calon yang mendapatkan suara terbanyak kedua. Jika beberapa calon mendapatkan suara terbanyak yang sama, ketua terpilih adalah dengan nomor peserta yang paling kecil dan wakilnya dengan nomor peserta terkecil berikutnya. 
//Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan data tidak valid. Data valid adalah bilangan bulat dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi berapa banyak suara yang valid.  Kemudian tercetak calon nomor berapa saja yang menjadi pasangan ketua RT dan wakil ketua RT yang baru. 
package main
import "fmt"
func main(){
	var suara int
	var totalSuara int
	var suaraValid int
	var suaraCalon [20]int
	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		totalSuara++
		if suara >= 1 && suara <= 20 {
			suaraValid++
			suaraCalon[suara-1]++
		}
	}
	fmt.Println(totalSuara)
	fmt.Println(suaraValid)
	var max1, max2 int
	var calonMax1, calonMax2 int
	for i := 0; i < 20; i++ {
		if suaraCalon[i] > max1 {
			max2 = max1
			calonMax2 = calonMax1
			max1 = suaraCalon[i]
			calonMax1 = i + 1
		} else if suaraCalon[i] > max2 {
			max2 = suaraCalon[i]
			calonMax2 = i + 1
		}
	}
	fmt.Printf("Ketua RT: Calon %d dengan %d suara\n", calonMax1, max1)
	fmt.Printf("Wakil Ketua RT: Calon %d dengan %d suara\n", calonMax2, max2)
}