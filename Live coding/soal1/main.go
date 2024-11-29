// Rizkulloh Alpriyansah
// 2311102142
package main

import (
	"fmt"
	"math"
)

func Bilangan_2311102142(bil int) (int, int) {
    digit := len(fmt.Sprintf("%d", bil))
    tengah := digit / 2

    if digit%2 == 0 {
        bil1 := bil / int(math.Pow(10, float64(tengah)))
        bil2 := bil % int(math.Pow(10, float64(tengah)))
        return bil1, bil2
    } else {
        bil1 := bil / int(math.Pow(10, float64(tengah)))
        bil2 := bil % int(math.Pow(10, float64(tengah)))
        return bil1, bil2
    }
}

func main() {
    var bil int

    fmt.Print("Masukkan bilangan bulat positif (>10): ")
    fmt.Scan(&bil)

    if bil <= 10 {
        fmt.Println("Bilangan harus lebih besar dari 10.")
    } else {
        bil1, bil2 := Bilangan_2311102142(bil)

        fmt.Println("Bilangan 1:", bil1)
        fmt.Println("Bilangan 2:", bil2)
        fmt.Println("Hasil penjumlahan:", bil1+bil2)
    }
    fmt.Println("")
}