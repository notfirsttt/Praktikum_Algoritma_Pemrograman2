package main

import (
	"fmt"
)

func printFactors(n_142, current_142 int) {
	if current_142 > n_142 {
		return
	}
	if n_142%current_142 == 0 {
		fmt.Print(current_142, " ")
	}
	printFactors(n_142, current_142+1)
}

func main() {
	var n_142 int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n_142)

	fmt.Println("Faktor dari", n_142, ":")
	printFactors(n_142, 1)
	fmt.Println()
}
