package main

import (
	"fmt"
) 

func selectionSort_142(arr_142 []int) { 
    n_142 := len(arr_142) 
    for i_142 := 0; i_142 < n_142-1; i_142++ { 
        minIdx_142 := i_142 
        for j_142 := i_142 + 1; j_142 < n_142; j_142++ { 
            if arr_142[j_142] < arr_142[minIdx_142] { 
                minIdx_142 = j_142 
            } 
        } 
        arr_142[i_142], arr_142[minIdx_142] = arr_142[minIdx_142], arr_142[i_142] 
    } 
} 

func main() { 
    var n_142 int 
    fmt.Print("Jumlah daerah kerabat Hercules tinggal: ") 
    fmt.Scan(&n_142) 

    if n_142 <= 0 || n_142 > 1000 { 
        return 
    } 

    daerah_142 := make([][]int, n_142) 

    for i_142 := 0; i_142 < n_142; i_142++ { 
        var m_142 int 
        fmt.Printf("\nJumlah kerabat di daerah %d: ", i_142+1) 
        fmt.Scan(&m_142) 

        if m_142 <= 0 || m_142 > 1000000 { 
            return 
        } 

        daerah_142[i_142] = make([]int, m_142) 
        fmt.Printf("Nomor rumah untuk daerah %d: ", i_142+1) 
        for j_142 := 0; j_142 < m_142; j_142++ { 
            fmt.Scan(&daerah_142[i_142][j_142]) 
        } 
        selectionSort_142(daerah_142[i_142]) 
    } 

    fmt.Println("\nHasil pengurutan nomor rumah:") 
    for i_142 := 0; i_142 < n_142; i_142++ { 
        fmt.Printf("Daerah %d: ", i_142+1) 
        for j_142 := 0; j_142 < len(daerah_142[i_142]); j_142++ { 
            fmt.Printf("%d ", daerah_142[i_142][j_142]) 
        } 
        fmt.Println() 
    } 
} 
