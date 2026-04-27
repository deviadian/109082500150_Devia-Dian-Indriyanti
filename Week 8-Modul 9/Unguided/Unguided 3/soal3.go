package main
import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var pemenang [100]string
	index := 0
	no := 1

	for {
		fmt.Printf("Pertandingan %d : ", no)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			fmt.Println("Hasil", no, ":", klubA)
			pemenang[index] = klubA
			index++
		} else if skorB > skorA {
			fmt.Println("Hasil", no, ":", klubB)
			pemenang[index] = klubB
			index++
		} else {
			fmt.Println("Hasil", no, ": Draw")
		}

		no++
	}

	fmt.Println("Pertandingan selesai")
}