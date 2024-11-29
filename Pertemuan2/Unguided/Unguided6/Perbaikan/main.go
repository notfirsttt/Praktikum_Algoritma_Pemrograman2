package main

import "fmt"

func main() {
	var nam_142 float64
	var nmk_142 string

	// Meminta input nilai
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam_142)

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam_142 > 80 {
		nmk_142 = "A"
	} else if nam_142 > 72.5 {
		nmk_142 = "AB"
	} else if nam_142 > 65 {
		nmk_142 = "B"
	} else if nam_142 > 57.5 {
		nmk_142 = "BC"
	} else if nam_142 > 50 {
		nmk_142 = "C"
	} else if nam_142 > 40 {
		nmk_142 = "D"
	} else {
		nmk_142 = "E"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Indeks untuk nilai akhir mata kuliah: %s\n", nmk_142)
}
