package main

import (
	"fmt"
)

const NMAX = 1000000

type Partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]Partai

func main() {
	var p tabPartai
	var jumlahPartai int

	jumlahPartai = inputSuara(&p)

	insertionSortDescending(&p, jumlahPartai)

	printHasil(p, jumlahPartai)
}

func inputSuara(p *tabPartai) int {
	var input, index, n int
	n = 0
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		index = posisi(*p, n, input)
		if index == -1 { 
			p[n].nama = input
			p[n].suara = 1
			n++
		} else { 
			p[index].suara++
		}
	}
	return n
}

func posisi(p tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if p[i].nama == nama {
			return i
		}
	}
	return -1 
}

func insertionSortDescending(p *tabPartai, n int) {
	var i, j int
	var temp Partai
	for i = 1; i < n; i++ {
		temp = p[i]
		j = i - 1
		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}
}

func printHasil(p tabPartai, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}
