package main

import "fmt"

// Definisikan tipe data set sebagai array dengan ukuran max 2022
type set [2022]int

// Fungsi untuk memeriksa apakah suatu nilai ada dalam set
func exist(T set, n int, val int) bool {
	var i int = 0
	var status bool = false
	for i < n && !status {
		status = T[i] == val
		i++
	}
	return status
}

// Fungsi untuk memasukkan data ke dalam set
func inputSet(T *set, n *int) {
	*n = 0
	var bilangan int
	fmt.Scan(&bilangan)
	for *n < 2022 && !exist(*T, *n, bilangan) {
		T[*n] = bilangan
		(*n)++
		fmt.Scan(&bilangan)
	}
}

// Fungsi untuk mencari irisan dari dua himpunan
func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	var j int = 0
	*h = 0
	for j < n {
		if exist(T2, m, T1[j]) {
			T3[*h] = T1[j]
			(*h)++
		}
		j++
	}
}

// Fungsi untuk mencetak set
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	// Input set pertama
	inputSet(&s1, &n1)
	// Input set kedua
	inputSet(&s2, &n2)

	// Mencari irisan kedua set
	findIntersection(s1, s2, n1, n2, &s3, &n3)

	// Menampilkan hasil irisan
	printSet(s3, n3)
}