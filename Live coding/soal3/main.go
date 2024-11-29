// Rizkulloh Alpriyansah
// 2311102142
package main

import "fmt"

func PerkalianRekursif_2311102142(n, m int) int {
    if m == 0 {
        return 0
    }
    return n + PerkalianRekursif_2311102142(n, m-1)
}

func main() {
    var n, m int

    fmt.Print("Masukkan bilangan n: ")
    fmt.Scan(&n)

    fmt.Print("Masukkan bilangan m: ")
    fmt.Scan(&m)

    Total := PerkalianRekursif_2311102142(n, m)

    fmt.Printf("Hasil dari %d x %d = %d ", n, m, Total)
}