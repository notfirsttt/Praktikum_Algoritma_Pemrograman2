package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
) 

func findMedian_142(numbers_142 []int) int { 
    sort.Ints(numbers_142) 
    n_142 := len(numbers_142) 
    if n_142%2 == 1 {  
        return numbers_142[n_142/2] 
    } 
    return (numbers_142[n_142/2-1] + numbers_142[n_142/2]) / 2 
} 

func main() { 
    reader_142 := bufio.NewReader(os.Stdin) 
    fmt.Println("Masukkan angka, gunakan 0 untuk menghitung median dan akhiri dengan -5313:") 

    input_142, _ := reader_142.ReadString('\n') 
    input_142 = strings.TrimSpace(input_142) 

    data_142 := strings.Split(input_142, " ") 
    var numbers_142 []int 

    for _, value_142 := range data_142 { 
        num_142, err_142 := strconv.Atoi(value_142) 
        if err_142 != nil { 
            fmt.Println("Input tidak valid, masukkan hanya angka.") 
            return 
        } 

        if num_142 == -5313 { 
            break  
        } else if num_142 == 0 { 
            if len(numbers_142) > 0 { 
                median_142 := findMedian_142(numbers_142) 
                fmt.Println(median_142)  
            } 
        } else { 
            numbers_142 = append(numbers_142, num_142) 
        } 
    } 
} 