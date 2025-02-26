package main

import (
	"fmt"
	"sort"
) 
 
const mMax = 7919 
 
type Buku struct { 
    judul    string 
    penulis  string 
    penerbit string 
    tahun    int 
    rating   float64 
} 
 
type DaftarBuku struct { 
    data   []Buku 
    nBuku  int 
} 
 
func main() { 
    var pustaka DaftarBuku 
    var cariRating float64 
 
    DaftarkanBuku(&pustaka) 
 
    CetakTerfavorit(&pustaka) 
 
    UrutBuku(&pustaka) 
 
    fmt.Println("\nDaftar buku setelah diurutkan berdasarkan rating:") 
    CetakSemuaBuku(&pustaka) 
 
    fmt.Println("\nMasukkan rating buku yang ingin dicari:") 
    fmt.Scan(&cariRating) 
    CariBuku(&pustaka, cariRating) 
} 
func DaftarkanBuku(pustaka *DaftarBuku) { 
    var n int 
    fmt.Print("Masukkan jumlah buku: ") 
    fmt.Scan(&n) 
 
    for i := 0; i < n; i++ { 
        var buku Buku 
        fmt.Printf("\nMasukkan data buku ke-%d:\n", i+1) 
        fmt.Print("Judul: ") 
        fmt.Scan(&buku.judul) 
        fmt.Print("Penulis: ") 
        fmt.Scan(&buku.penulis) 
        fmt.Print("Penerbit: ") 
        fmt.Scan(&buku.penerbit) 
        fmt.Print("Tahun: ") 
        fmt.Scan(&buku.tahun) 
        fmt.Print("Rating: ") 
        fmt.Scan(&buku.rating) 
 
        pustaka.data = append(pustaka.data, buku) 
    } 
    pustaka.nBuku = n 
} 
 
func CetakTerfavorit(pustaka *DaftarBuku) { 
    if pustaka.nBuku == 0 { 
        fmt.Println("Tidak ada data buku.") 
        return 
    } 
 
    tertinggi := pustaka.data[0] 
    for _, buku := range pustaka.data { 
        if buku.rating > tertinggi.rating { 
            tertinggi = buku 
        } 
    } 
 
    fmt.Println("\nBuku dengan rating tertinggi:") 
    fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %.2f\n", 
        tertinggi.judul, tertinggi.penulis, 
tertinggi.penerbit, tertinggi.tahun, tertinggi.rating) 
} 
 
func UrutBuku(pustaka *DaftarBuku) {
	sort.Slice(pustaka.data, func(i, j int) bool { 
        return pustaka.data[i].rating > 
pustaka.data[j].rating 
    }) 
} 
 
func CetakSemuaBuku(pustaka *DaftarBuku) { 
    if pustaka.nBuku == 0 { 
        fmt.Println("Tidak ada data buku.") 
        return 
    } 
 
    for _, buku := range pustaka.data { 
        fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %.2f\n", 
            buku.judul, buku.penulis, buku.penerbit, 
buku.tahun, buku.rating) 
    } 
} 
 
func CariBuku(pustaka *DaftarBuku, rating float64) { 
    for _, buku := range pustaka.data { 
        if buku.rating == rating { 
            fmt.Println("\nBuku yang ditemukan:") 
            fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %.2f\n", 
                buku.judul, buku.penulis, buku.penerbit, 
buku.tahun, buku.rating) 
            return 
        } 
    } 
 
    fmt.Println("\nBuku dengan rating tersebut tidak ditemukan.") 
} 