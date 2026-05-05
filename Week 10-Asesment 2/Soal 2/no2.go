package main
import "fmt"

const nmax = 51

type mahasiswa struct {
	nim   string
	nama  string
	nilai int
}

type arrMahasiswa [nmax]mahasiswa

func inputData(A *arrMahasiswa, n *int) {
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&A[i].nim, &A[i].nama, &A[i].nilai)
	}
}

func cariNilaiPertama(A arrMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if A[i].nim == nim {
			return A[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(A arrMahasiswa, n int, nim string) int {
	max := -1
	for i := 0; i < n; i++ {
		if A[i].nim == nim {
			if A[i].nilai > max {
				max = A[i].nilai
			}
		}
	}
	return max
}

func main() {
	var A arrMahasiswa
	var n int
	var nimCari string

	inputData(&A, &n)

	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(A, n, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(A, n, nimCari)

	fmt.Println("NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya :", nimCari)
	fmt.Println("Nilai pertama dari NIM", nimCari, "adalah", nilaiPertama)
	fmt.Println("Nilai terbesar dari NIM", nimCari, "adalah", nilaiTerbesar)
}