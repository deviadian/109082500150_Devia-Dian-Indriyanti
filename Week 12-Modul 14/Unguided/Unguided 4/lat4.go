package main

import "fmt"

func main() {
	var daerah int
	fmt.Scan(&daerah)

	for i := 0; i < daerah; i++ {

		var n int
		fmt.Scan(&n)

		data := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scan(&data[j])
		}

		var ganjil []int
		var genap []int

		for j := 0; j < n; j++ {
			if data[j]%2 == 1 {
				ganjil = append(ganjil, data[j])
			} else {
				genap = append(genap, data[j])
			}
		}

		for j := 0; j < len(ganjil)-1; j++ {
			for k := j + 1; k < len(ganjil); k++ {
				if ganjil[j] > ganjil[k] {
					ganjil[j], ganjil[k] = ganjil[k], ganjil[j]
				}
			}
		}

		for j := 0; j < len(genap)-1; j++ {
			for k := j + 1; k < len(genap); k++ {
				if genap[j] < genap[k] {
					genap[j], genap[k] = genap[k], genap[j]
				}
			}
		}

		for j := 0; j < len(ganjil); j++ {
			fmt.Print(ganjil[j], " ")
		}
		fmt.Println()

		for j := 0; j < len(genap); j++ {
			fmt.Print(genap[j], " ")
		}
		fmt.Println()
	}
}