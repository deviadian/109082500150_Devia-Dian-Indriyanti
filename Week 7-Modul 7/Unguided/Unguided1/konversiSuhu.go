package main
import "fmt"

type suhu float64
func CelciusToReamur(celcius suhu) suhu {
	return(4.0 / 5.0) * celcius
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return(9.0 / 5.0) * celcius + 32
}

func CelciusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("masukkan suhu (celcius): ")
	fmt.Scan(&c)

	r := CelciusToReamur(c)
	f := CelciusToFahrenheit(c)
	k := CelciusToKelvin(c)

	fmt.Printf("\n%.0f celcius = %.1f reamur\n", c, r)
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", c, f)
	fmt.Printf("%.0f celcius = %.2f kelvin\n", c, k)
}