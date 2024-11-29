package main

import (
	"fmt"
)

func printOdd(current_142, n_142 int) {
	if current_142 > n_142 {
		return
	}
	fmt.Print(current_142, " ")
	printOdd(current_142+2, n_142)
}

func main() {
	var n_142 int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n_142)

	printOdd(1, n_142)
	fmt.Println()
}