package main

import (
	"fmt"
)

// Fungsi untuk menghitung nilai hampiran akar 2 berdasarkan iterasi K
func calculateSqrt2(k_142 int) float64 {
	product_142 := 1.0 // Mulai dengan nilai produk 1

	// Iterasi dari 0 hingga K
	for i_142 := 0; i_142 <= k_142; i_142++ {
		numerator_142 := (4*float64(i_142) + 2) * (4*float64(i_142) + 2) // (4k + 2)^2
		denominator_142 := (4*float64(i_142) + 1) * (4*float64(i_142) + 3) // (4k + 1)(4k + 3)
		product_142 *= numerator_142 / denominator_142
	}

	return product_142 // Mengembalikan nilai hasil perkalian
}

func main() {
	var k_142 int

	// Membaca input nilai K
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k_142)

	// Menghitung nilai hampiran sqrt(2)
	result_142 := calculateSqrt2(k_142)

	// Menampilkan hasil dengan 10 angka di belakang koma
	fmt.Printf("Nilaif(10) = %.10f\n", result_142)
}