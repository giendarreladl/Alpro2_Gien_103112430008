package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&c)

	r := (4.0 / 5.0) * c
	f := (9.0/5.0)*c + 32
	k := c + 273

	fmt.Printf("Derajat Reamur: %.0f\n", r)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", f)
	fmt.Printf("Derajat Kelvin: %.0f\n", k)
}
