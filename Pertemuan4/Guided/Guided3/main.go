package main

import "fmt"

func volumeTabung(jari_jari, tinggi int) float64 {
var luasAlas, volume float64
luasAlas = 3.14 * float64(jari_jari*jari_jari)
volume = luasAlas * float64(tinggi)
return volume
}
func main() {
var r, t int
var v1, v2 float64
r = 5
t = 10
v1 = volumeTabung(r, t)
v2 = volumeTabung(r, t) + volumeTabung(15, t)
fmt.Println("volume tabung pertama:", v1)
fmt.Println("volume tabung kedua (gabungan):", v2)
fmt.Println("volume tabung dengan jari-jari 14 dan tinggi 100:", volumeTabung(14, 100))
}