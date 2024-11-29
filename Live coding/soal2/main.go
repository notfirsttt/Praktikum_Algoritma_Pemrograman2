// Rizkulloh Alpriyansah
// 2311102142
package main

import (
	"fmt"
	"strconv"
)

func Hadiah_2311102142(nomor int) string {
    angka := strconv.Itoa(nomor)
    Angka := make(map[rune]bool)
    AngkaSamaSemua := true
    AngkaBerulang := false

    for i, angkaSekarang := range angka {
        if i > 0 && angkaSekarang != rune(angka[0]) {
            AngkaSamaSemua = false
        }
        if Angka[angkaSekarang] {
            AngkaBerulang = true
        }
        Angka[angkaSekarang] = true
    }

    if AngkaSamaSemua {
        return "Hadiah A"
    }
    if AngkaBerulang {
        return "Hadiah C"
    }
    return "Hadiah B"
}

func main() {
    var jumlahPeserta int
    fmt.Print("Masukkan jumlah peserta: ")
    fmt.Scan(&jumlahPeserta)

    var hadiahA, hadiahB, hadiahC int

    for i := 1; i <= jumlahPeserta; i++ {
        var nomorKartu int
        fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
        fmt.Scan(&nomorKartu)

        hadiah := Hadiah_2311102142(nomorKartu)
        fmt.Println(hadiah)

        switch hadiah {
        case "Hadiah A":
            hadiahA++
        case "Hadiah B":
            hadiahB++
        case "Hadiah C":
            hadiahC++
        }
    }

    fmt.Printf("\nJumlah yang memperoleh Hadiah A: %d\n", hadiahA)
    fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
    fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC)
}