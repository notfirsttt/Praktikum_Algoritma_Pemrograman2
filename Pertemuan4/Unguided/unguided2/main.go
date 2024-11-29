package main

import (
	"fmt"
	"strings"
)

type Peserta struct {
	nama       string
	totalSoal  int
	totalWaktu int
}

func hitungSkor(soal []int) (int, int) {
	totalSoal := 0
	totalWaktu := 0

	for _, waktu := range soal {
		totalWaktu += waktu
	}

	for _, waktu := range soal {
		if waktu > 0 { 
			totalSoal++
		}
	}

	if totalWaktu > 300 {
		return 0, 0
	}

	return totalSoal, totalWaktu
}

func main() {
	var nama string
	var waktuTemp int
	var soal []int

	var pemenang Peserta
	maxSoal := 0

	for {
		fmt.Scan(&nama)

		if strings.ToLower(nama) == "selesai" {
			break
		}

		soal = make([]int, 8) 

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktuTemp)
			soal[i] = waktuTemp
		}

		totalSoal, totalWaktu := hitungSkor(soal)

		if totalSoal > maxSoal || (totalSoal == maxSoal && totalWaktu < pemenang.totalWaktu) {
			pemenang = Peserta{nama, totalSoal, totalWaktu}
			maxSoal = totalSoal
		}
	}

	if pemenang.totalSoal > 0 {
		fmt.Printf("%s %d %d\n", pemenang.nama, pemenang.totalSoal, pemenang.totalWaktu)
	} else {
		fmt.Println("Tidak ada peserta yang valid.")
	}
}
