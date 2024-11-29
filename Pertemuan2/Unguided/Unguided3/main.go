package main

import (
	"fmt"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		// Menerima input berat kedua kantong dalam satu baris
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		// Cek jika salah satu kantong mencapai atau lebih dari 9 kg
		if beratKiri >= 9 || beratKanan >= 9 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat antara kantong kiri dan kanan
		selisih := beratKiri - beratKanan
		if selisih < 0 {
			selisih = -selisih
		}

		// Cek apakah selisih berat melebihi 9 kg
		if selisih > 9 {
			fmt.Println("Selisih berat antara kantong kiri dan kanan melebihi 9 kg.")
		}
	}
}