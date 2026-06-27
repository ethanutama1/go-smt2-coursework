//Pada pemilihan ketua RT yang baru saja berlangsung, terdapat 20 calon ketua yang bertanding memperebutkan suara warga. Perhitungan suara dapat segera dilakukan karena warga cukup mengisi formulir dengan nomor dari calon ketua RT yang dipilihnya. Seperti biasa, selalu ada pengisian yang tidak tepat atau dengan nomor pilihan di luar yang tersedia, sehingga data juga harus divalidasi. Tugas Anda untuk membuat program mencari siapa yang memenangkan pemilihan ketua RT. 
//Buatlah program pilkart yang akan membaca, memvalidasi, dan menghitung suara yang diberikan dalam pemilihan ketua RT tersebut. Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan data tidak valid. Data valid adalah integer dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. 
//Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi berapa banyak suara yang valid. Kemudian sejumlah baris yang mencetak data para calon apa saja yang mendapatkan suara. 

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
	for i := 0; i < 20; i++ {
		if suaraCalon[i] > 0 {
			fmt.Printf("Calon %d: %d suara\n", i+1, suaraCalon[i])
		}
	}
}