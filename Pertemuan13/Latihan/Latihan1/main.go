package main

import "fmt"

func main() {
	var x, y, z []int

	exists := false

	fmt.Printf("Masukkan elemen untuk array x: ")
	for !exists {
		var n int
		fmt.Scan(&n)

		for _, v := range x {
			if v == n {
				exists = true
			}
		}

		if !exists {
			x = append(x, n)
		}
	}

	exists = false

	fmt.Printf("Masukkan elemen untuk array y: ")
	for !exists {
		var n int
		fmt.Scan(&n)

		for _, v := range y {
			if v == n {
				exists = true
			}
		}

		if !exists {
			y = append(y, n)
		}
	}

	for _, v := range x {
		for _, w := range y {
			if v == w {
				z = append(z, v)
			}
		}
	}

	fmt.Println("Elemen yang sama:", z)
}