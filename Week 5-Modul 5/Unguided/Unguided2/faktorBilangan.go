package main
import "fmt"

func faktorRekursif(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorRekursif(n, i+1)
}

func main() {
	var n int
	fmt.Print("masukkan N: ")
	fmt.Scan(&n)
	fmt.Print("faktor: ")
	faktorRekursif(n, 1)
}