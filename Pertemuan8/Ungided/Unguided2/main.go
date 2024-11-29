package main

import (
	"fmt"
	"math"
)

func main() {
	var n_142 int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n_142)

	array_142 := make([]int, n_142)
	fmt.Println("Masukkan elemen array:")
	for i_142 := 0; i_142 < n_142; i_142++ {
		fmt.Printf("Elemen ke-%d: ", i_142)
		fmt.Scan(&array_142[i_142])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan keseluruhan isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Hitung rata-rata elemen array")
		fmt.Println("7. Hitung simpangan baku elemen array")
		fmt.Println("8. Hitung frekuensi bilangan tertentu")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan_142 int
		fmt.Scan(&pilihan_142)

		switch pilihan_142 {
		case 1:
			fmt.Println("Isi array:")
			for _, val_142 := range array_142 {
				fmt.Print(val_142, " ")
			}
			fmt.Println()

		case 2:
			fmt.Println("Elemen dengan indeks ganjil:")
			for i_142 := 1; i_142 < len(array_142); i_142 += 2 {
				fmt.Print(array_142[i_142], " ")
			}
			fmt.Println()

		case 3:
			fmt.Println("Elemen dengan indeks genap:")
			for i_142 := 0; i_142 < len(array_142); i_142 += 2 {
				fmt.Print(array_142[i_142], " ")
			}
			fmt.Println()

		case 4:
			var x_142 int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x_142)
			fmt.Print("Elemen dengan indeks kelipatan ", x_142, ": ")
			if x_142 > 0 {
				for i_142 := x_142; i_142 < len(array_142); i_142 += x_142 {
					fmt.Print(array_142[i_142], " ")
				}
			} else {
				fmt.Println("Nilai x harus lebih besar dari 0.")
			}
			fmt.Println()

		case 5:
			var idx_142 int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&idx_142)
			if idx_142 >= 0 && idx_142 < len(array_142) {
				array_142 = append(array_142[:idx_142], array_142[idx_142+1:]...)
				fmt.Println("Array setelah penghapusan:")
				for _, val_142 := range array_142 {
					fmt.Print(val_142, " ")
				}
				fmt.Println()
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case 6:
			sum_142 := 0
			for _, val_142 := range array_142 {
				sum_142 += val_142
			}
			rataRata_142 := float64(sum_142) / float64(len(array_142))
			fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata_142)

		case 7:
			sum_142 := 0
			for _, val_142 := range array_142 {
				sum_142 += val_142
			}
			mean_142 := float64(sum_142) / float64(len(array_142))

			var variance_142 float64
			for _, val_142 := range array_142 {
				variance_142 += math.Pow(float64(val_142)-mean_142, 2)
			}
			variance_142 /= float64(len(array_142))
			stdDev_142 := math.Sqrt(variance_142)
			fmt.Printf("Simpangan baku elemen array: %.2f\n", stdDev_142)

		case 8:
			var target_142 int
			fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
			fmt.Scan(&target_142)
			count_142 := 0
			for _, val_142 := range array_142 {
				if val_142 == target_142 {
					count_142++
				}
			}
			fmt.Printf("Frekuensi bilangan %d: %d kali\n", target_142, count_142)

		case 9:
			// Keluar dari program
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}