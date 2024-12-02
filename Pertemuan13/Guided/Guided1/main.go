package main

import "fmt"

//Insertionsort untuk mengurutkan array secara discanding
func InsertionSort(arr []int){
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		//Geser elemen yang lebih kecil ke kanan
		for j > 0 && temp >arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j+1] = temp
	}		
}

func main(){
	data := []int{5, 2, 9, 1, 4, 6}
	fmt.Println("Data sebelum di urutkan:", data)
	InsertionSort(data)
	fmt.Println("Data setelah di urutkan:", data)
}