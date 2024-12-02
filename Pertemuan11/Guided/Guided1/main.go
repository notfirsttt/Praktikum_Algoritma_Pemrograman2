package main

import "fmt"

// Fungsi untuk melakukan Selection Sort
func selectionSort(arr []int) { 
    n := len(arr) 
    for i := 0; i < n-1; i++ { // Iterasi dari indeks pertama sampai kedua terakhir 
        idxMin := i // Anggap elemen pada indeks i adalah yang terkecil 
 
        // Cari elemen terkecil di sisa array 
        for j := i + 1; j < n; j++ { 
            if arr[j] < arr[idxMin] { 
                idxMin = j // Update indeks nilai terkecil 
            } 
        } 
 
        // Tukar elemen terkecil dengan elemen di posisi i 
        if idxMin != i { 
            temp := arr[i] 
            arr[i] = arr[idxMin] 
            arr[idxMin] = temp 
        } 
    } 
} 
 
func main() { 
    // Data yang akan diurutkan 
    data := []int{64, 25, 12, 22, 11} 
 
    fmt.Println("Data sebelum diurutkan:", data) 
 
    // Panggil fungsi Selection Sort 
    selectionSort(data) 
 
    fmt.Println("Data setelah diurutkan:", data) 
} 
