package main
import "fmt"

func main() {
	var x int
	var suara [21] int
	suaraMasuk := 0
	suaraSah := 0

	for {
		fmt.Scan(&x)
		suaraMasuk++
		if x == 0 {
			break
		}
		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
	}

	ketua := 1

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		} else if suara[i] == suara[ketua] && i < ketua {
			ketua = i
		}
	}

	wakil := -1

	for i := 1; i <= 20; i++ {
		if i == ketua {
			continue
		}
		if wakil == -1 {
			wakil = i
		}
		if suara[i] > suara[wakil] {
			wakil = i
		} else if suara[i] == suara[wakil] && i < wakil {
			wakil = i
		}
	}
	fmt.Println("Suara masuk: ", suaraMasuk)
	fmt.Println("Suara sah: ", suaraSah)
	fmt.Println("Ketua RT: ", ketua)
	fmt.Println("Wakil ketua: ", wakil)
}