package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. tampil semua
	fmt.Println("\nIsi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	// b. indeks ganjil
	fmt.Println("\n\nIndeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// c. indeks genap
	fmt.Println("\n\nIndeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// d. kelipatan x
	var x int
	fmt.Print("\n\nMasukkan x: ")
	fmt.Scan(&x)

	fmt.Println("Indeks kelipatan x:")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// e. hapus elemen (manual, tanpa append biar basic)
	var idx int
	fmt.Print("\n\nIndeks yang dihapus: ")
	fmt.Scan(&idx)

	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Array setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	// f. rata-rata
	total := 0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	rata := float64(total) / float64(n)
	fmt.Println("\n\nRata-rata:", rata)

	// h. frekuensi
	var cari int
	fmt.Print("\nMasukkan angka yang dicari: ")
	fmt.Scan(&cari)

	freq := 0
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}