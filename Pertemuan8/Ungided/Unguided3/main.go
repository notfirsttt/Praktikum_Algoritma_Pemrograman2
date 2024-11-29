package main

import (
	"fmt"
)

func main() {
	var klubA_142, klubB_142 string
	var skorA_142, skorB_142 int
	var hasil_142 []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA_142)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB_142)

	pertandingan_142 := 1
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan_142)
		fmt.Scan(&skorA_142, &skorB_142)

		if skorA_142 < 0 || skorB_142 < 0 {
			break
		}

		if skorA_142 > skorB_142 {
			hasil_142 = append(hasil_142, fmt.Sprintf("Hasil %d: %s", pertandingan_142, klubA_142))
		} else if skorB_142 > skorA_142 {
			hasil_142 = append(hasil_142, fmt.Sprintf("Hasil %d: %s", pertandingan_142, klubB_142))
		} else {
			hasil_142 = append(hasil_142, fmt.Sprintf("Hasil %d: Draw", pertandingan_142))
		}

		pertandingan_142++
	}

	fmt.Println("\nDaftar Hasil Pertandingan:")
	for _, h_142 := range hasil_142 {
		fmt.Println(h_142)
	}

	fmt.Println("Pertandingan selesai")
}
