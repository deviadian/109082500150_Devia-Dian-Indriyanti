package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	max := pustaka[0]

	for i := 1; i < n; i++ {
		if pustaka[i].rating > max.rating {
			max = pustaka[i]
		}
	}

	fmt.Println(
		max.judul,
		max.penulis,
		max.penerbit,
		max.tahun,
	)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var pass, i int
	var temp Buku

	for pass = 1; pass < n; pass++ {
		temp = pustaka[pass]
		i = pass

		for i > 0 && temp.rating > pustaka[i-1].rating {
			pustaka[i] = pustaka[i-1]
			i--
		}

		pustaka[i] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5

	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Println(
			pustaka[i].judul,
			pustaka[i].rating,
		)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	low := 0
	high := n - 1
	found := false

	for low <= high && !found {
		mid := (low + high) / 2

		if pustaka[mid].rating == r {
			found = true

			fmt.Println(
				pustaka[mid].judul,
				pustaka[mid].penulis,
				pustaka[mid].penerbit,
				pustaka[mid].tahun,
				pustaka[mid].eksemplar,
				pustaka[mid].rating,
			)

		} else if r > pustaka[mid].rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}