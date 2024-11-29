package main

import "fmt"

func main() {
	cetak2(5)
}
func cetak2(x int) {
	fmt.Println(x)
	cetak2(x + 1)
}