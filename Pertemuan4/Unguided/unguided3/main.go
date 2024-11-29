package main

import "fmt"

// Fungsi untuk mencetak deret bilangan
func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(n) // Cetak angka terakhir (1)
}

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Masukkan bilangan positif kurang dari 1.000.000")
	}
}
