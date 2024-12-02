package main

import "fmt" 
 
type arrInt [4321]int 
 
func selectionSort1(T *arrInt, n int) { 
    var t, i, j, idx_min int 
 
    i = 1 
    for i <= n-1 { 
        idx_min = i - 1 
        j = i 
        for j < n { 
            if T[idx_min] > T[j] { 
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
    var data arrInt 
    var n int 
 
    fmt.Print("Masukkan jumlah elemen array: ") 
    fmt.Scan(&n) 
 
    fmt.Println("Masukkan elemen array:") 
    for i := 0; i < n; i++ { 
        fmt.Scan(&data[i]) 
    } 
 
    fmt.Println("\nData sebelum diurutkan:") 
    for i := 0; i < n; i++ { 
        fmt.Printf("%d ", data[i]) 
} 
	selectionSort1(&data, n) 
	fmt.Println("\n\nData setelah diurutkan:") 
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", data[i])
	} 
} 