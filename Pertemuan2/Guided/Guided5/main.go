// By Rizkulloh NIM 2311102142
package main

import (
	"fmt"
)

func main() {
	// Inisialisasi array untuk menampung 5 integer dan 3 karakter
	var nums [5]int
	var chars [3]rune

	// Meminta input untuk 5 data integer
	fmt.Println("Masukkan 5 angka integer (32-127):")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	// Meminta input untuk 3 karakter
	fmt.Println("Masukkan 3 karakter:")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
		// Mengabaikan newline dari input sebelumnya
		if chars[i] == '\n' {
			i--
		}
	}

	// Output pertama: 5 integer yang dikonversi menjadi karakter ASCII
	fmt.Println("Keluaran:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", nums[i])
	}
	fmt.Println()

	// Output kedua: 3 karakter yang diinputkan
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", chars[i])
	}
	fmt.Println()
}