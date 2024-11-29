package main

import "fmt"

func main() {
	cetak(5)
}
func cetak(x int) {
	if x == 10 {
		fmt.Println(x) //base case
	} else {
		fmt.Println(x)
		cetak(x + 1)
	}
}