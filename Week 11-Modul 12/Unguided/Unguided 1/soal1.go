package main
import "fmt"

func sequentialSearch(target int) bool {
	for i := 1; i <= 20; i++ {
		if i == target {
			return true
		}
	}
	return false
}

func main() {
	var x int
	total := 0
	sah := 0

	var suara [21]int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		total++

		if sequentialSearch(x) {
			sah++
			suara[x]++
		}
	}

	fmt.Printf("Suara masuk: %d\n\n", total)
	fmt.Printf("Suara sah: %d\n\n", sah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}