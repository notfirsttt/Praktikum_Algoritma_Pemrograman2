// Rizkulloh Alpriyansah 2311102142
package main

import "fmt"

func cetakFibo(n int) {
	var f1, f2, f3 int
	f2 = 0
	f3 = 1

	for i := 1; i <= n; i++ {
		fmt.Println(f3)
		f1 = f2
		f2 = f3
		f3 = f1 + f2

	}
}
func main() {
	cetakFibo(10)
}