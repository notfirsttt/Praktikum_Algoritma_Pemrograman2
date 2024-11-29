// By Rizkulloh 2311102142
package main

import (
	"fmt"
	"math"
)

// Fungsi buat hitung jarak antara dua titik (x1, y1) dan (x2, y2)
func menghitungJarak(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

// Ngecek apakah titik (x, y) ada di dalam lingkaran dengan pusat (px, py) dan jari-jari r
func diDalamLingkaran(px, py, r, x, y float64) bool {
    return menghitungJarak(px, py, x, y) <= r
}

func main() {
    var intiX1, intiY1, radius1 int
    var intiX2, intiY2, radius2 int
    var titikX, titikY int

    fmt.Println("Masukkan koordinat pusat dan jari-jari lingkaran 1 (px1, py1, r1): ")
    fmt.Scan(&intiX1, &intiY1, &radius1)

    fmt.Println("Masukkan koordinat pusat dan jari-jari lingkaran 2 (px2, py2, r2): ")
    fmt.Scan(&intiX2, &intiY2, &radius2)

    fmt.Println("Masukkan koordinat titik sembarang (x, y): ")
    fmt.Scan(&titikX, &titikY)

    diLingkaranKe1 := diDalamLingkaran(float64(intiX1), float64(intiY1), float64(radius1), float64(titikX), float64(titikY))
    diLingkaranKe2 := diDalamLingkaran(float64(intiX2), float64(intiY2), float64(radius2), float64(titikX), float64(titikY))

    if diLingkaranKe1 && diLingkaranKe2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if diLingkaranKe1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if diLingkaranKe2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
    fmt.Println()
}