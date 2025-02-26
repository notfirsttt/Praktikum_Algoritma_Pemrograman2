package main

import "fmt"

func main() {
	var data_142 []int
	var input_142 int

	for {
		_, err_142 := fmt.Scan(&input_142)
		if err_142 != nil {
			break
		}

		if input_142 == -5313541 {
			return
		} else if input_142 == 0 {
			if len(data_142) == 0 {
				continue
			}
			selectionSort_142(data_142)
			n_142 := len(data_142)
			if n_142%2 == 1 {
				fmt.Println(data_142[n_142/2])
			} else {
				fmt.Println((data_142[n_142/2-1] + data_142[n_142/2]) / 2)
			}
		} else {
			data_142 = append(data_142, input_142)
		}
	}
}

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
