// By Rizkulloh 2311102142
package main

import "fmt"

func f(x int) int {
    return x * x
}

func g(x int) int {
    return x - 2
}

func h(x int) int {
    return x + 1
}

func fogog(x int) int {
    return f(g(h(x)))
}

func gohof(x int) int {
    return g(h(f(x)))
}

func hofog(x int) int {
    return h(f(g(x)))
}

func main() {
    var a, b, c int
    fmt.Print("Masukkan tiga angka dipisahkan oleh spasi: \n")
    fmt.Scan(&a, &b, &c)

    fmt.Println("fogog(", a, ")=", fogog(a))
    fmt.Println("gohof(", b, ")=", gohof(b))
    fmt.Println("hofog(", c, ")=", hofog(c))
	fmt.Println()
}