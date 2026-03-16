package main
import "fmt"

func f(x int) int{
	return x * x
}

func g(x int) int{
	return x - 2
}

func h(x int) int{
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("masukkan nilai x: ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai y: ")
	fmt.Scan(&b)
	fmt.Print("masukkan nilai z: ")
	fmt.Scan(&c)

	fmt.Println("F(G(H(", a, "))) : ", f(g(h(a))))
	fmt.Println("G(H(F(", b, "))) : ", g(h(f(b))))
	fmt.Println("H(F(G(", c, "))) : ", h(f(g(c))))
}