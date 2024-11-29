package main

import (
	"fmt"
)

func hitungBiayaKirim(berat_142 int) int {
	// Menghitung berat dalam kilogram dan gram
	kg_142 := berat_142 / 1000
	gram_142 := berat_142 % 1000

	// Biaya pengiriman per kilogram
	biayaPerKg_142 := 10000
	biayaTotal_142 := kg_142 * biayaPerKg_142

	// Biaya tambahan untuk sisa gram
	biayaTambahan_142 := 0
	if kg_142 >= 10 {
		biayaTambahan_142 = 0 // Gratis biaya tambahan jika berat lebih dari 10 kg
	} else {
		if gram_142 >= 500 {
			biayaTambahan_142 = gram_142 * 5 // Rp. 5 per gram jika sisa >= 500 gram
		} else {
			biayaTambahan_142 = gram_142 * 15 // Rp. 15 per gram jika sisa < 500 gram
		}
	}

	// Total biaya
	return biayaTotal_142 + biayaTambahan_142
}

func main() {
	var berat_142 int

	// Meminta input berat dari pengguna
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat_142)

	// Menghitung berat dalam kg dan gram
	kg_142 := berat_142 / 1000
	gram_142 := berat_142 % 1000

	// Menghitung total biaya pengiriman
	biayaPerKg_142 := 10000 * kg_142
	biayaTambahan_142 := 0

	// Kondisi untuk biaya tambahan berdasarkan sisa berat gram
	if kg_142 >= 10 {
		biayaTambahan_142 = 0 // Gratis biaya tambahan jika lebih dari 10 kg
	} else {
		if gram_142 >= 500 {
			biayaTambahan_142 = gram_142 * 5 // Rp. 5 per gram untuk sisa >= 500 gram
		} else {
			biayaTambahan_142 = gram_142 * 15 // Rp. 15 per gram untuk sisa < 500 gram
		}
	}

	// Menghitung total biaya
	totalBiaya_142 := biayaPerKg_142 + biayaTambahan_142

	// Menampilkan hasil
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg_142, gram_142)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg_142, biayaTambahan_142)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya_142)
}