package main
import "fmt"

const NMAX = 100000

var data [NMAX] int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2

		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}