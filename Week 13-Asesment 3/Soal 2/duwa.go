package main
import "fmt"

type Pemain struct {
	nama   string
	gol    int
	assist int
}

func main() {
	var n int
	fmt.Scan(&n)

	var pemain [1001]Pemain

	for i := 0; i < n; i++ {
		fmt.Scan(&pemain[i].nama, &pemain[i].gol, &pemain[i].assist)
	}

	for i := 1; i < n; i++ {
		temp := pemain[i]
		j := i - 1

		for j >= 0 && (pemain[j].gol < temp.gol ||
			(pemain[j].gol == temp.gol && pemain[j].assist < temp.assist)) {

			pemain[j+1] = pemain[j]
			j--
		}

		pemain[j+1] = temp
	}

	for i := 0; i < n; i++ {
		fmt.Println(pemain[i].nama, pemain[i].gol, pemain[i].assist)
	}
}