package main

import (
	"fmt"
)

func main() {
	// Input suhu dalam Celsius
	var celsius float64
	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scan(&celsius)

	// Konversi ke Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Konversi ke Reamur
	reamur := celsius * 4 / 5

	// Konversi ke Kelvin
	kelvin := celsius + 273.15

	// Output hasil konversi
	fmt.Printf("Temperatur Celsius: %.2f\n", celsius)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}