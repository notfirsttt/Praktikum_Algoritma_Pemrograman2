package main

import (
	"fmt"
)

func power(x_142, y_142 int) int {
	if y_142 == 0 {
		return 1
	}
	return x_142 * power(x_142, y_142-1)
}

func main() {
	var x_142, y_142 int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x_142)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y_142)

	result_142 := power(x_142, y_142)
	fmt.Println("Hasil:", result_142)
}
