package main
import "fmt"

type angka int
type kata string

type Buku struct {
	judul kata
	penulis kata
	penerbit kata
	tahunTerbit angka
	jumlahHalaman angka
}

func main() {
	var b Buku
	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("masukkan judul buku: ")
	fmt.Scan(&b.judul)

	fmt.Print("masukkan nama penulis: ")
	fmt.Scan(&b.penulis)

	fmt.Print("masukkan penerbit: ")
	fmt.Scan(&b.penerbit)

	fmt.Print("masukkan tahun terbit: ")
	fmt.Scan(&b.tahunTerbit)

	fmt.Print("masukkan jumlah halaman: ")
	fmt.Scan(&b.jumlahHalaman)

	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("judul buku: ", b.judul)
	fmt.Println("penulis: ", b.penulis)
	fmt.Println("penerbit: ", b.penerbit)
	fmt.Println("tahun terbit: ", b.tahunTerbit)
	fmt.Println("jumlah halaman: ", b.jumlahHalaman)
}