package main

import "fmt" 
 
type mahasiswa struct { 
    nama, nim, kelas, jurusan string 
    ipk                       float64 
} 
 
type arrMhs [2023]mahasiswa 
 
func selectionSort2(T *arrMhs, n int) { 
 
    var i, j, idx_min int 
    var t mahasiswa 
 
    i = 1 
    for i <= n-1 { 
        idx_min = i - 1 
        j = i 
        for j < n { 
            if T[idx_min].ipk > T[j].ipk { 
                idx_min = j 
            } 
            j = j + 1 
        } 
        t = T[idx_min] 
        T[idx_min] = T[i-1] 
        T[i-1] = t 
        i = i + 1 
    } 
} 
 
func main() { 
    var data arrMhs 
    var n int 
 
    fmt.Print("Masukkan jumlah mahasiswa: ") 
    fmt.Scan(&n) 
 
    fmt.Println("Masukkan data mahasiswa (Nama, NIM, Kelas, Jurusan, IPK):") 
    for i := 0; i < n; i++ { 
        fmt.Printf("Data mahasiswa ke-%d:\n", i+1) 
        fmt.Print("Nama: ") 
        fmt.Scan(&data[i].nama) 
        fmt.Print("NIM: ") 
        fmt.Scan(&data[i].nim) 
        fmt.Print("Kelas: ") 
        fmt.Scan(&data[i].kelas) 
        fmt.Print("Jurusan: ") 
        fmt.Scan(&data[i].jurusan) 
        fmt.Print("IPK: ") 
        fmt.Scan(&data[i].ipk) 
    } 
 
    fmt.Println("\nData mahasiswa sebelum diurutkan:") 
    for i := 0; i < n; i++ { 
        fmt.Printf("%s (%s) - %s - %s - IPK: %.2f\n", 
data[i].nama, data[i].nim, data[i].kelas, data[i].jurusan, 
data[i].ipk) 
    } 
    selectionSort2(&data, n) 
 
    fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan IPK (terkecil-terbesar):") 
    for i := 0; i < n; i++ { 
        fmt.Printf("%s (%s) - %s - %s - IPK: %.2f\n", 
data[i].nama, data[i].nim, data[i].kelas, data[i].jurusan, 
data[i].ipk) 
    } 
}