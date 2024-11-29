// Rizkulloh Alpriyansah
// 2311102142

package main

import "fmt" 

func main() { 
    var beratAnak_142 [1000]float64 

    var jumlahAnak_142 int 
    fmt.Print("Masukkan jumlah anak kelinci: ") 
    fmt.Scan(&jumlahAnak_142) 

    fmt.Println("Masukkan berat masing-masing anak kelinci:") 
    for i := 0; i < jumlahAnak_142; i++ { 
        fmt.Scan(&beratAnak_142[i]) 
    } 

    beratTerkecil_142 := beratAnak_142[0] 
    beratTerbesar_142 := beratAnak_142[0] 

    for i := 1; i < jumlahAnak_142; i++ { 
        if beratAnak_142[i] < beratTerkecil_142 { 
            beratTerkecil_142 = beratAnak_142[i] 
        } 
        if beratAnak_142[i] > beratTerbesar_142 { 
            beratTerbesar_142 = beratAnak_142[i] 
        } 
    } 

    fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil_142) 
    fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar_142) 
} 