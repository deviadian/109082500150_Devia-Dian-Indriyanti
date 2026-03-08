package main
import "fmt"

func main() {
	var target, tabungan, total, hari int

	fmt.Print("masukkan target uang yg ingin dicapai: ")
	fmt.Scanln(&target)

	total = 0
	hari = 0

	for total < target {
		hari++
		fmt.Printf("masukkan nominal tabungan hari ke-%d: ", hari)
		fmt.Scanln(&tabungan)

		total = total + tabungan
	}

	fmt.Printf("selamat target anda tercapai dalam %d hari.\n", hari)
	fmt.Printf("total tabungan anda terkumpul: Rp%d\n", total)
}