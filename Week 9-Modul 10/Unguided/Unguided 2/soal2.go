package main
import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	berat := make([]float64, x)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	hasil := make([]float64, 0)
	totalAll := 0.0

	for i := 0; i < x; i += y {
		total := 0.0
		
		for j := i; j < i+y && j < x; j++ {
			total += berat[j]
		}
		hasil = append(hasil, total)
		totalAll += total
	}

	for i := 0; i < len(hasil); i++ {
		fmt.Printf("%.1f ", hasil[i])
	}
	fmt.Println()

	rata := totalAll / float64(len(hasil))
	fmt.Printf("%.1f\n", rata)
}