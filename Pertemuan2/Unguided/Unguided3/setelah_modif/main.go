package main

import (
	"fmt"
)

func main() {
	var beratKantongKiri_142, beratKantongKanan_142 float64
	for {
		// Menerima input berat kedua kantong dalam satu baris
		fmt.Print("Masukan berat belanjaan di kedua kantong (kg): ")
		fmt.Scan(&beratKantongKiri_142, &beratKantongKanan_142)

		// Cek jika salah satu kantong memiliki berat negatif
		if beratKantongKiri_142 < 0 || beratKantongKanan_142 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek jika total berat kedua kantong melebihi 150 kg
		totalBerat_142 := beratKantongKiri_142 + beratKantongKanan_142
		if totalBerat_142 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat antara kantong kiri dan kanan
		selisihBerat_142 := beratKantongKiri_142 - beratKantongKanan_142
		if selisihBerat_142 < 0 {
			selisihBerat_142 = -selisihBerat_142
		}

		// Menampilkan hasil apakah sepeda motor akan oleng atau tidak
		if selisihBerat_142 >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}