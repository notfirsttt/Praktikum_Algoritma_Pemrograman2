package main

import (
	"fmt"
)

func printDescending(n_142 int) {
	if n_142 < 1 {
		return
	}
	fmt.Print(n_142, " ")
	printDescending(n_142 - 1)
}

func printAscending(current_142, n_142 int) {
	if current_142 > n_142 {
		return
	}
	fmt.Print(current_142, " ")
	printAscending(current_142 + 1, n_142)
}

func main() {
	var n_142 int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n_142)
	printDescending(n_142)
	printAscending(2, n_142)
	fmt.Println()
}
