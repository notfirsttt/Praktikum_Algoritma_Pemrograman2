package main

import "fmt"

const nProv int = 34

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	var i int
	for i = 0; i < nProv; i++ {
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	var idx int = 0
	var i int
	for i = 1; i < nProv; i++ {
		if tumbuh[idx] < tumbuh[i] {
			idx = i
		}
	}
	return idx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var i int
	var result float64
	for i = 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			result = (1 + tumbuh[i]) * float64(pop[i])
			fmt.Println(prov[i], result)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	var found int = -1
	var i int = 0
	for i < nProv && found == -1 {
		if prov[i] == nama {
			found = i
		}
		i++
	}
	return found
}

func main() {
	var TProvinsi NamaProv
	var TPopulasi PopProv
	var TPertumbuhan TumbuhProv
	var cari string
	var idxTercepat, idxProvinsi int

	InputData(&TProvinsi, &TPopulasi, &TPertumbuhan)

	fmt.Scan(&cari)

	idxTercepat = ProvinsiTercepat(TPertumbuhan)
	fmt.Println(TProvinsi[idxTercepat])

	idxProvinsi = IndeksProvinsi(TProvinsi, cari)
	if idxProvinsi != -1 {
		fmt.Println(idxProvinsi)
	} else {
		fmt.Println("Provinsi tidak ditemukan.")
	}

	Prediksi(TProvinsi, TPopulasi, TPertumbuhan)
}