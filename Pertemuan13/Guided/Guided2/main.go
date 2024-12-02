package main

import "fmt"

// Struct mahasiswa

type mahasiswa struct {
    nama, nim, kelas, jurusan string
    ipk float64
}

// InsertionSort untuk mengurutkan array mahasiswa berdasarkan nama (descending)

func insertionSort(arr []mahasiswa) {
    n := len(arr)
    for i := 1; i < n; i++ {
        temp := arr[i]
        j := i
        // Geser elemen yang lebih kecil (berdasarkan nama) ke kanan
        for j > 0 && temp.nama > arr[j-1].nama {
            arr[j] = arr[j-1]
            j--
        }
        arr[j] = temp
    }
}

func main() {

    data := []mahasiswa{
        {"Kyla", "101", "IF1", "Teknik Informatika", 4.0},
        {"Azzahra", "102", "IF2", "Teknik Informatika", 3.8},
        {"Kinan", "103", "IF1", "Teknik Informatika", 3.9},
        {"Ara", "104", "IF2", "Teknik Informatika", 3.4},
    }
    
    fmt.Println("Data sebelum diurutkan:")
    for _, m := range data {
        fmt.Println(m)
    }
    
    insertionSort(data)
    
    fmt.Println("\nData setelah diurutkan (descending berdasarkan nama):")
    for _, m := range data {
        fmt.Println(m)
    }
}