package main

import "fmt"

func main() {
	var arr [100]int
	var n, x int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		arr[n] = x
		n++
	}

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	selisih := arr[1] - arr[0]
	tetap := true

	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != selisih {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}