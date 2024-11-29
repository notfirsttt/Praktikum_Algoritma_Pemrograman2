package main

import (
	"fmt"
) 

func calculateMinMax_142(weights_142 []float64, minWeight_142, maxWeight_142 *float64) { 
    *minWeight_142 = weights_142[0] 
    *maxWeight_142 = weights_142[0] 
    for _, weight_142 := range weights_142 { 
        if weight_142 < *minWeight_142 { 
            *minWeight_142 = weight_142 
        } 
        if weight_142 > *maxWeight_142 { 
            *maxWeight_142 = weight_142 
        } 
    } 
} 

func averageWeight_142(weights_142 []float64) float64 { 
    totalWeight_142 := 0.0 
    for _, weight_142 := range weights_142 { 
        totalWeight_142 += weight_142 
    } 
    return totalWeight_142 / float64(len(weights_142)) 
} 

func main() { 
    var totalData_142 int 
    fmt.Print("Masukan banyak data berat balita: ") 
    fmt.Scan(&totalData_142) 

    weights_142 := make([]float64, totalData_142) 
    for i_142 := 0; i_142 < totalData_142; i_142++ { 
        fmt.Printf("Masukan berat balita ke-%d: ", i_142+1) 
        fmt.Scan(&weights_142[i_142]) 
    } 

    var minWeight_142, maxWeight_142 float64 
    calculateMinMax_142(weights_142, &minWeight_142, &maxWeight_142) 

    avgWeight_142 := averageWeight_142(weights_142) 
    fmt.Printf("Berat balita minimum: %.2f kg\n", minWeight_142) 
    fmt.Printf("Berat balita maksimum: %.2f kg\n", maxWeight_142) 
    fmt.Printf("Rata-rata berat balita: %.2f kg\n", avgWeight_142) 
}