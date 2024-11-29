package main

import (
	"fmt"
)

func main() {
	var number_142 int

	// Input bilangan bulat
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number_142)

	// Mencari dan menampilkan semua faktor dari bilangan
	fmt.Printf("Faktor: ")
	factors_142 := findFactors(number_142)

	// Menampilkan faktor-faktor
	for _, factor_142 := range factors_142 {
		fmt.Printf("%d ", factor_142)
	}
	fmt.Println()

	// Mengecek apakah bilangan tersebut prima
	if len(factors_142) == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}

// Fungsi untuk mencari faktor dari bilangan
func findFactors(n_142 int) []int {
	var factors_142 []int
	for i_142 := 1; i_142 <= n_142; i_142++ {
		if n_142%i_142 == 0 {
			factors_142 = append(factors_142, i_142)
		}
	}
	return factors_142
}