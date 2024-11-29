package main

import (
	"fmt"
)

// Fungsi untuk menghitung nilai hampiran akar 2 berdasarkan iterasi K
func hitungAkar2(iterasi_142 int) float64 {
	hasilPerkalian_142 := 1.0 // Mulai dengan nilai produk 1

	// Iterasi dari 0 hingga iterasi yang dimasukkan
	for indeks_142 := 0; indeks_142 < iterasi_142; indeks_142++ { // Menggunakan < bukan <=
		pembilang_142 := (4*float64(indeks_142) + 2) * (4*float64(indeks_142) + 2) // (4k + 2)^2
		penyebut_142 := (4*float64(indeks_142) + 1) * (4*float64(indeks_142) + 3) // (4k + 1)(4k + 3)
		hasilPerkalian_142 *= pembilang_142 / penyebut_142
	}
	return hasilPerkalian_142 // Mengembalikan nilai hasil perkalian
}

func main() {
	var jumlahIterasi_142 int

	// Membaca input nilai jumlah iterasi
	fmt.Print("Masukkan jumlah iterasi K: ")
	fmt.Scan(&jumlahIterasi_142)

	// Memastikan jumlah iterasi tidak negatif
	if jumlahIterasi_142 < 0 {
		fmt.Println("Jumlah iterasi K tidak boleh negatif.")
		return
	}

	// Menghitung nilai hampiran sqrt(2)
	hasilAkar2_142 := hitungAkar2(jumlahIterasi_142)

	// Menampilkan hasil dengan 10 angka di belakang koma
	fmt.Printf("Nilai hampiran akar 2 = %.10f\n", hasilAkar2_142)
}
