package main
import "fmt"

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func didalam(c Lingkaran, p Titik) bool {
	dx := p.x - c.pusat.x
	dy := p.y - c.pusat.y
	return dx*dx+dy*dy <= c.r*c.r
}

func main() {
	var c1, c2 Lingkaran
	var p Titik

	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.r)
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.r)
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}