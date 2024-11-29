package main

import (
	"fmt"
)

func printStars(n_142, current_142 int) {
	if current_142 > n_142 {
		return
	}
	for i := 0; i < current_142; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printStars(n_142, current_142+1)
}

func main() {
	var n_142 int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n_142)
	printStars(n_142, 1)
}
