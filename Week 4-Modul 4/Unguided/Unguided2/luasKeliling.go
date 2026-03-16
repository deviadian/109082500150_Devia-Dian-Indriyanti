package main
import "fmt"

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Println()
	fmt.Println("masukkan sisi: ", sisi)
	fmt.Println("luas persegi: ", luas)
	fmt.Println("keliling persegi: ", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println()
	fmt.Println("masukkan panjang: ", panjang)
	fmt.Println("masukkan lebar: ", lebar)
	fmt.Println("luas persegi panjang: ", luas)
	fmt.Println("keliling persegi panjang: ", keliling)
}

func hitungLingkaran(jarijari float64) {
	luas := 3.14 * jarijari * jarijari
	keliling := 2 * 3.14 * jarijari

	fmt.Println()
	fmt.Println("masukkan jari-jari: ", jarijari)
	fmt.Println("luas lingkaran: ", luas)
	fmt.Println("keliling lingkaran: ", keliling)
}

func main() {
	var pilihan int
	var sisi, panjang, lebar int
	var jarijari float64

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("masukkan sisi: ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		fmt.Print("masukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("masukkan lebar: ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		fmt.Print("masukkan jari-jari: ")
		fmt.Scan(jarijari)
		hitungLingkaran(jarijari)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}