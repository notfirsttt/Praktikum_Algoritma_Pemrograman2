package main

import (
	"fmt"
)

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&n)

	results := make([][]int, n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah %d: ", i+1)
		fmt.Scan(&m)

		rumahKerabat := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat di daerah %d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumahKerabat[j])
		}

		var ganjil, genap []int
		for _, rumah := range rumahKerabat {
			if rumah%2 == 1 {
				ganjil = append(ganjil, rumah)
			} else {
				genap = append(genap, rumah)
			}
		}

		fmt.Println("Ganjil sebelum sorting:", ganjil)
		fmt.Println("Genap sebelum sorting:", genap)

		selectionSortDesc(ganjil)
		selectionSortAsc(genap)

		fmt.Println("Ganjil setelah sorting:", ganjil)
		fmt.Println("Genap setelah sorting:", genap)

		results[i] = append(ganjil, genap...)
	}

	fmt.Println("\nHasil:")
	for i, result := range results {
		fmt.Printf("Daerah %d: ", i+1)
		for _, rumah := range result {
			fmt.Printf("%d ", rumah)
		}
		fmt.Println()
	}
}