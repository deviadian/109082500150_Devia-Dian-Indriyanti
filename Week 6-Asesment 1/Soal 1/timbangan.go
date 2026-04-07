package main
import "fmt"
const pi float64 = 3.14

func volume (r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		selisih := m1 - m2
		if selisih < 0 {
			selisih = -selisih
		}
		fmt.Println(selisih)
	}
}

func main() {
	var r float64
	var t1, t2 float64
	var p1, p2 float64

	fmt.Print("masukkan jari-jari alas tabung: ")
	fmt.Scan(&r)
	fmt.Print("masukkan tinggi zat cair kiri: ")
	fmt.Scan(&t1)
	fmt.Print("masukkan massa jenis kiri: ")
	fmt.Scan(&p1)
	fmt.Print("masukkan tinggi zat cair kanan: ")
	fmt.Scan(&t2)
	fmt.Print("masukkan massa jenis kanan: ")
	fmt.Scan(&p2)

	m1 := massa(r, t1, p1)
	m2 := massa(r, t2, p2)

	display(m1, m2)
}