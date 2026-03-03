package main
import "fmt"

func main() {
	var berat, tinggi, bmi float64

	fmt.Print("masukkan berat badan (kg): ")
	fmt.Scanln(&berat)
	fmt.Print("masukkan tinggi badan (m): ")
	fmt.Scanln(&tinggi)

	bmi = berat / (tinggi * tinggi)

	fmt.Printf("nilai bmi anda adalah: %.2f\n", bmi)
}