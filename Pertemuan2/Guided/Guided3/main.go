package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari int
	fmt.Print("Jejari: ")
	fmt.Scanf("%d", &jejari)

	volume := hitungVolumeBola(float64(jejari))
	luas := hitungLuasKulitBola(float64(jejari))

	fmt.Println("Bola dengan jejari", jejari, "memiliki volume", volume, "dan luas kulit", luas)
}

func hitungVolumeBola(jejari float64) float64 {
	const pi = 3.1415926535
	return (4.0 / 3.0) * pi * math.Pow(jejari, 3)
}

func hitungLuasKulitBola(jejari float64) float64 {
	const pi = 3.1415926535
	return 4 * pi * math.Pow(jejari, 2)
}