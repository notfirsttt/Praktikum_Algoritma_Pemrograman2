package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1) 
}

func permutasi(n, r int) int {
	if n >= r {
		return faktorial(n) / faktorial(n-r) 
	}
	return 0
}

func kombinasi(n, r int) int {
	if n >= r {
		return faktorial(n) / (faktorial(r) * faktorial(n-r))
	}
	return 0
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Printf("%d %d\n", permutasi(a, c), kombinasi(a, c))
	fmt.Printf("%d %d\n", permutasi(b, d), kombinasi(b, d))
}
