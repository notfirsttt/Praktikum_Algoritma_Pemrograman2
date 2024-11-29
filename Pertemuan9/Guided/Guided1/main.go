// Rizkulloh Alpriyansah
// 2311102142

package main

import "fmt" 
 
type mahasiswa struct { 
    nama, nim, kelas, jurusan string 
    ipk                       float64 
} 
 
type arrMhs [2023]mahasiswa 
 
func IPK_1(T arrMhs, n int) float64 { 
    var terkecil float64 = T[0].ipk 
    var j int = 1 
 
    for j < n { 
        if terkecil > T[j].ipk { 
            terkecil = T[j].ipk 
        } 
        j = j + 1 
    } 
    return terkecil 
} 
 
func main() { 
    var mahasiswaArr arrMhs 
    mahasiswaArr[0] = mahasiswa{"Alice", "2211102123", "IF1", "Teknik Informatika", 3.2} 
    mahasiswaArr[1] = mahasiswa{"Bob", "2211102124", "IF-2", "Teknik Informatika", 3.8} 
    mahasiswaArr[2] = mahasiswa{"Charlie", "22111021235", "IF-1", "Teknik Informatika", 2.9} 
    mahasiswaArr[3] = mahasiswa{"Diana", "22111022189", "IF3", "Teknik Informatika", 3.0} 
	
	n := 4 
 
    ipkTerkecil := IPK_1(mahasiswaArr, n) 
 
    fmt.Printf("IPK terkecil adalah: %.2f\n", ipkTerkecil) 
} 