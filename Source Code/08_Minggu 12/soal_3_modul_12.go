package main
import "fmt"

func binarySearch(arr []int, k int) int{
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if arr[mid] == k {
			return mid
		}
		if arr[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
func main(){
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	result := binarySearch(arr, k)
	if result != -1 {
		fmt.Printf("Angka %d ditemukan pada indeks %d\n", k, result)
	} else {
		fmt.Printf("Angka %d tidak ditemukan dalam array\n", k)
	}
}