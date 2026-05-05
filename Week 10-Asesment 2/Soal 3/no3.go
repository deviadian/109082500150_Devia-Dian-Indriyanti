package main
import "fmt"

type Provinsi struct {
	nama       string
	populasi   int
	pertumbuhan float64
}

const NMAX int = 10

type TabelProvinsi [NMAX]Provinsi

func main() {
	var daftar TabelProvinsi
	var cariNama string
	
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	
	for i := 0; i < NMAX; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&daftar[i].nama, &daftar[i].populasi, &daftar[i].pertumbuhan)
	}

	fmt.Scan(&cariNama)

	tercepat := cariPertumbuhanTercepat(daftar)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", tercepat)

	idx := cariIndexProvinsi(daftar, cariNama)
	if idx != -1 {
		fmt.Printf("Data provinsi yang dicari : %s\n", daftar[idx].nama)
	}

	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	tampilkanPrediksi(daftar)
}

func cariPertumbuhanTercepat(T TabelProvinsi) string {
	maxIdx := 0
	for i := 1; i < NMAX; i++ {
		if T[i].pertumbuhan > T[maxIdx].pertumbuhan {
			maxIdx = i
		}
	}
	return T[maxIdx].nama
}

func cariIndexProvinsi(T TabelProvinsi, nama string) int {
	for i := 0; i < NMAX; i++ {
		if T[i].nama == nama {
			return i
		}
	}
	return -1
}

func tampilkanPrediksi(T TabelProvinsi) {
	for i := 0; i < NMAX; i++ {
		if T[i].pertumbuhan > 0.02 {
			prediksi := float64(T[i].populasi) * (1 + T[i].pertumbuhan)
			fmt.Printf("%s %.0f\n", T[i].nama, prediksi)
		}
	}
}