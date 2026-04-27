package main
import "fmt"

const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune
	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var balik tabel

	for i := 0; i < n; i++ {
		balik[i] = t[i]
	}

	balikanArray(&balik, n)

	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	balikanArray(&tab, m)
	cetakArray(tab, m)
}