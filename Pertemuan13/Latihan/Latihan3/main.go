package main

import "fmt"

const nProv int = 34

// Tipe data array untuk menyimpan nama provinsi, populasi, dan angka pertumbuhan
type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

// Fungsi untuk memasukkan data provinsi, populasi, dan angka pertumbuhan
func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	// Mengisi array prov, pop, dan tumbuh dengan data
	var i int
	for i = 0; i < nProv; i++ {
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

// Fungsi untuk mencari provinsi dengan pertumbuhan tercepat
func ProvinsiTercepat(tumbuh TumbuhProv) int {
	var idx int = 0
	var i int
	// Mencari indeks dengan pertumbuhan tercepat
	for i = 1; i < nProv; i++ {
		if tumbuh[idx] < tumbuh[i] {
			idx = i
		}
	}
	return idx
}

// Fungsi untuk menampilkan prediksi jumlah penduduk provinsi dengan pertumbuhan > 2%
func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var i int
	var result float64
	// Menampilkan prediksi jumlah penduduk
	for i = 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			result = (1 + tumbuh[i]) * float64(pop[i])
			fmt.Println(prov[i], result)
		}
	}
}

// Fungsi untuk mencari indeks provinsi berdasarkan nama
func IndeksProvinsi(prov NamaProv, nama string) int {
	var found int = -1
	var i int = 0
	// Mencari indeks provinsi berdasarkan nama
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

	// Memasukkan data provinsi
	InputData(&TProvinsi, &TPopulasi, &TPertumbuhan)

	// Menerima nama provinsi yang ingin dicari
	fmt.Scan(&cari)

	// Mencari provinsi dengan pertumbuhan tercepat dan menampilkan nama provinsi
	idxTercepat = ProvinsiTercepat(TPertumbuhan)
	fmt.Println(TProvinsi[idxTercepat])

	// Mencari indeks provinsi yang dicari dan menampilkan nama provinsi
	idxProvinsi = IndeksProvinsi(TProvinsi, cari)
	if idxProvinsi != -1 {
		fmt.Println(idxProvinsi)
	} else {
		fmt.Println("Provinsi tidak ditemukan.")
	}

	// Menampilkan prediksi jumlah penduduk provinsi dengan pertumbuhan lebih dari 2%
	Prediksi(TProvinsi, TPopulasi, TPertumbuhan)
}