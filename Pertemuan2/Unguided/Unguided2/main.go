package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Membuat reader untuk membaca input dari pengguna
	reader_142 := bufio.NewReader(os.Stdin)

	// Meminta input jumlah bunga yang akan dimasukkan (bilangan bulat positif N)
	fmt.Print("N: ")
	var N_142 int
	for {
		// Baca input dari pengguna
		input_142, err := reader_142.ReadString('\n')
		if err != nil {
			fmt.Printf("Error membaca input: %v\n", err)
			return
		}

		// Konversi input ke integer
		N_142, err = strconv.Atoi(strings.TrimSpace(input_142))
		if err != nil || N_142 <= 0 {
			fmt.Println("Harap masukkan bilangan bulat positif.")
		} else {
			break
		}
	}

	// Inisialisasi variabel pita (string) untuk menyimpan nama bunga
	var pita_142 string
	var count_142 int // Menyimpan jumlah bunga yang dimasukkan

	// Loop untuk menerima input nama bunga sebanyak N kali
	for i_142 := 1; i_142 <= N_142; i_142++ {
		fmt.Printf("Bunga %d: ", i_142) // Menambahkan instruksi

		// Membaca input dari pengguna
		input_142, err := reader_142.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return // Keluar dari program jika ada kesalahan
		}

		// Menghapus spasi dan karakter newline dari input
		input_142 = strings.TrimSpace(input_142)

		// Cek jika pengguna mengetik "SELESAI"
		if strings.ToUpper(input_142) == "SELESAI" {
			break // Menghentikan input jika "SELESAI" dimasukkan
		}

		// Menggabungkan nama bunga dengan pita menggunakan " – " sebagai pemisah
		if pita_142 == "" {
			pita_142 = input_142 // Jika pita masih kosong, langsung masukkan nama bunga
		} else {
			pita_142 = pita_142 + " – " + input_142 // Jika sudah ada isinya, tambahkan dengan pemisah " – "
		}

		count_142++ // Menambah jumlah bunga yang dimasukkan
	}

	// Menampilkan isi pita dan jumlah bunga setelah semua input dimasukkan
	fmt.Println("Pita:", pita_142)
	fmt.Printf("Bunga: %d\n", count_142)
}
