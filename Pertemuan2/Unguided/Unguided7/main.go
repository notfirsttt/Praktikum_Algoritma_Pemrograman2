package main

import (
	"fmt"
)

func main() {
	var b_142 int

	// Menerima input dari pengguna
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&b_142)

	// Memeriksa apakah b lebih besar dari 0
	if b_142 <= 0 {
		fmt.Println("Bilangan harus lebih besar dari 0.")
		return
	}

	fmt.Printf("Bilangan: %d\nFaktor: ", b_142)

	factors_142 := []int{} // Untuk menyimpan faktor-faktor

	// Mencari faktor dari b
	for i_142 := 1; i_142 <= b_142; i_142++ {
		if b_142%i_142 == 0 {
			factors_142 = append(factors_142, i_142)
			if i_142 == b_142 {
				fmt.Print(i_142)
			} else {
				fmt.Print(i_142, ", ")
			}
		}
	}

	// Menentukan apakah b adalah bilangan prima
	isPrime_142 := len(factors_142) == 2

	// Menampilkan hasil
	fmt.Printf("\nPrima: %t\n", isPrime_142)
}