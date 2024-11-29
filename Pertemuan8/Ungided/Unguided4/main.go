package main

import (
	"fmt"
)

const NMAX_142 = 127

type tabel_142 [NMAX_142]rune

func isiArray_142(t_142 *tabel_142, n_142 *int) {
	*n_142 = 0
	fmt.Println("Masukkan teks (akhiri dengan '.'): ")
	for {
		var c_142 rune
		fmt.Scanf("%c", &c_142)
		if c_142 == '.' {
			break
		}
		t_142[*n_142] = c_142
		(*n_142)++
	}
}

func cetakArray_142(t_142 tabel_142, n_142 int) {
	for i_142 := 0; i_142 < n_142; i_142++ {
		fmt.Printf("%c", t_142[i_142])
	}
	fmt.Println()
}

func balikkanArray_142(t_142 *tabel_142, n_142 int) {
	for i_142, j_142 := 0, n_142-1; i_142 < j_142; i_142, j_142 = i_142+1, j_142-1 {
		t_142[i_142], t_142[j_142] = t_142[j_142], t_142[i_142]
	}
}

func palindrom_142(t_142 tabel_142, n_142 int) bool {
	for i_142 := 0; i_142 < n_142/2; i_142++ {
		if t_142[i_142] != t_142[n_142-i_142-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab_142 tabel_142
	var n_142 int

	isiArray_142(&tab_142, &n_142)

	fmt.Print("Teks: ")
	cetakArray_142(tab_142, n_142)

	balikkanArray_142(&tab_142, n_142)
	fmt.Print("Reverse teks: ")
	cetakArray_142(tab_142, n_142)

	if palindrom_142(tab_142, n_142) {
		fmt.Println("Palindrom: True")
	} else {
		fmt.Println("Palindrom: False")
	}
}