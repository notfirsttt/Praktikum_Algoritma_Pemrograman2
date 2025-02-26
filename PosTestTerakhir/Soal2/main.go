package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputMhs(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(N)

	var i int = 0
	for i < *N {
		fmt.Printf("Masukkan NIM, Nama, dan Nilai mahasiswa ke-%d: ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
		i++
	}
}

func findFirstScore(T arrayMahasiswa, N int, nim string) int {
	var i int
	for i = 0; i < N; i++ {
		if T[i].NIM == nim {
			return i 
		}
	}
	return -1 
}

func findMaxScore(T arrayMahasiswa, N int, nim string) int {
	found := findFirstScore(T, N, nim)
	if found != -1 {
		idxMax := found
		for i := found + 1; i < N; i++ {
			if T[i].NIM == nim && T[i].nilai > T[idxMax].nilai {
				idxMax = i
			}
		}
		return idxMax
	}
	return found
}

func main() {
	var A arrayMahasiswa
	var M int
	var NIM string

	inputMhs(&A, &M)

	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&NIM)

	idx1 := findFirstScore(A, M, NIM)
	if idx1 == -1 {
		fmt.Println("NIM", NIM, "tidak ditemukan.")
	} else {
		idx2 := findMaxScore(A, M, NIM)
		fmt.Println("Nilai pertama dari NIM", NIM, "adalah", A[idx1].nilai)
		fmt.Println("Nilai terbesar dari NIM", NIM, "adalah", A[idx2].nilai)
	}
}