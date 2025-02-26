package main

import (
	"fmt"
	"sort"
)

func main() {
	var data []int

	fmt.Println("Masukkan")
	for {
		var num int
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		data = append(data, num)
	}

	if len(data) == 0 {
		fmt.Println("Tidak ada data yang dimasukkan.")
		return
	}

	sort.Ints(data)

	fmt.Println("Keluaran")
	for _, val := range data {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	status := checkDistance(data)
	fmt.Println(status)
}

func checkDistance(arr []int) string {
	if len(arr) < 2 {
		return "Data berjarak tidak tetap"
	}

	distance := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != distance {
			return "Data berjarak tidak tetap"
		}
	}

	return fmt.Sprintf("Data berjarak %d", distance)
}
