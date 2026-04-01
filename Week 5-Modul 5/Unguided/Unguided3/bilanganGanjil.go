package main
import "fmt"

func tampilGanjil(i, n int) {
	if i > n {
		return
	}
	fmt.Print(i, " ")
	tampilGanjil(i+2, n)
}

func main() {
	var n int
	fmt.Print("masukkan N: ")
	fmt.Scan(&n)
	tampilGanjil(1, n)
}