package main

import "fmt" 

func main() { 
    var totalFish_142, fishPerContainer_142 int 
    fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ") 
    fmt.Scan(&totalFish_142, &fishPerContainer_142) 

    fishWeights_142 := make([]float64, totalFish_142) 
    fmt.Println("Masukkan berat masing-masing ikan:") 
    for i := 0; i < totalFish_142; i++ { 
        fmt.Scan(&fishWeights_142[i]) 
    } 

    totalContainers_142 := (totalFish_142 + fishPerContainer_142 - 1) / fishPerContainer_142  
    containerWeights_142 := make([]float64, totalContainers_142) 
    for i := 0; i < totalFish_142; i++ { 
        containerIdx_142 := i / fishPerContainer_142 
        containerWeights_142[containerIdx_142] += fishWeights_142[i] 
    } 

    fmt.Println("Total berat di setiap wadah:") 
    for i := 0; i < totalContainers_142; i++ { 
        fmt.Printf("Wadah %d: %.2f kg\n", i+1, containerWeights_142[i]) 
    } 

    fmt.Println("Berat rata-rata per wadah:") 
    for i := 0; i < totalContainers_142; i++ { 
        fishInContainer_142 := fishPerContainer_142 
        if i == totalContainers_142-1 && totalFish_142%fishPerContainer_142 != 0 { 
            fishInContainer_142 = totalFish_142 % fishPerContainer_142 
        } 
        averageWeight_142 := containerWeights_142[i] / float64(fishInContainer_142) 
        fmt.Printf("Wadah %d: %.2f kg\n", i+1, averageWeight_142) 
    } 
} 