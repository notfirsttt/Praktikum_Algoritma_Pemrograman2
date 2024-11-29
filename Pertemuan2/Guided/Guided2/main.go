package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	kabisat := cekKabisat(tahun)
	fmt.Println("Kabisat: ", kabisat)
}

func cekKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%4 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	}
	return false
}