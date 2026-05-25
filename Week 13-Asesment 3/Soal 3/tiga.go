package main
import "fmt"

const NMAX = 100000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	var i int

	for i = 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}

	return -1
}

func main() {
	var p tabPartai
	var n int
	var x int
	var idx int
	var i, j int
	var temp partai

	n = 0

	for {
		fmt.Scan(&x)

		if x == -1 {
			break
		}

		idx = posisi(p, n, x)

		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
	}

	for i = 1; i < n; i++ {
		temp = p[i]
		j = i - 1

		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}

		p[j+1] = temp
	}

	for i = 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
}