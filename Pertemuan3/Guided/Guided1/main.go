// By RIzkulloh 2311102142
package main

import "fmt"

func hitungFaktorial(angka int) int {
	if angka == 0 || angka == 1 {
		return 1
	}
	hasil := 1
	for i := 2; i <= angka; i++ {
		hasil *= i
	}
	return hasil
}

func hitungPermutasi(total, ambil int) int {
	return hitungFaktorial(total) / hitungFaktorial(total-ambil)
}

func hitungKombinasi(total, ambil int) int {
	return hitungFaktorial(total) / (hitungFaktorial(ambil) * hitungFaktorial(total-ambil))
}

func main() {
	var nomer1, nomer2, nomer3, nomer4 int
	var Permutasi1, Permutasi2 int
	var Kombinasi1, Kombinasi2 int

	fmt.Println("Masukkan (angka1, angka2, angka3, angka4):")
	fmt.Scan(&nomer1, &nomer2, &nomer3, &nomer4)

	Permutasi1 = hitungPermutasi(nomer1, nomer3)
	Kombinasi1 = hitungKombinasi(nomer1, nomer3)

	Permutasi2 = hitungPermutasi(nomer2, nomer4)
	Kombinasi2 = hitungKombinasi(nomer2, nomer4)

	fmt.Println("Masukan\t|\tKeluaran")
	fmt.Println(Permutasi1, Kombinasi1,"\t|\t",Permutasi2, Kombinasi2)
}