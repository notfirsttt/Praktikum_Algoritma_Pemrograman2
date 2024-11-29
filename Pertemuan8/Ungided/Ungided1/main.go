package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x_142, y_142 int
}

type Lingkaran struct {
	titikPusat_142 Titik
	radius_142     int
}

func jarak_142(a_142, b_142 Titik) float64 {
	return math.Sqrt(float64((a_142.x_142-b_142.x_142)*(a_142.x_142-b_142.x_142) + (a_142.y_142-b_142.y_142)*(a_142.y_142-b_142.y_142)))
}

func dalamLingkaran_142(t_142 Titik, l_142 Lingkaran) bool {
	return jarak_142(t_142, l_142.titikPusat_142) <= float64(l_142.radius_142)
}

func main() {
	var cx1_142, cy1_142, r1_142 int
	var cx2_142, cy2_142, r2_142 int
	var px_142, py_142 int

	fmt.Println("Koordinat titik pusat & radius lingkaran 1 (cx cy r):")
	fmt.Scan(&cx1_142, &cy1_142, &r1_142)

	fmt.Println("Koordinat titik pusat & radius lingkaran 2 (cx cy r):")
	fmt.Scan(&cx2_142, &cy2_142, &r2_142)

	fmt.Println("Koordinat titik sembarang (x y):")
	fmt.Scan(&px_142, &py_142)

	lingkaran1_142 := Lingkaran{Titik{cx1_142, cy1_142}, r1_142}
	lingkaran2_142 := Lingkaran{Titik{cx2_142, cy2_142}, r2_142}
	titik_142 := Titik{px_142, py_142}

	inLingkaran1_142 := dalamLingkaran_142(titik_142, lingkaran1_142)
	inLingkaran2_142 := dalamLingkaran_142(titik_142, lingkaran2_142)

	if inLingkaran1_142 && inLingkaran2_142 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1_142 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2_142 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}