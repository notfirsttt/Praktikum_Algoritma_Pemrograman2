package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Urutan warna yang benar
	correctOrder_142 := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input untuk 5 percobaan
	reader_142 := bufio.NewReader(os.Stdin)
	success_142 := true

	for i_142 := 1; i_142 <= 5; i_142++ {
		fmt.Printf("Percobaan %d: ", i_142)

		// Membaca input dari pengguna
		input_142, _ := reader_142.ReadString('\n')
		input_142 = strings.TrimSpace(input_142)

		// Memisahkan input berdasarkan spasi
		colors_142 := strings.Split(input_142, " ")

		// Mengecek apakah urutan warna sesuai
		for j_142 := 0; j_142 < 4; j_142++ {
			if colors_142[j_142] != correctOrder_142[j_142] {
				success_142 = false
				break
			}
		}

		// Jika ada percobaan yang tidak sesuai, keluar dari loop
		if !success_142 {
			break
		}
	}

	// Menampilkan hasil
	if success_142 {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}
