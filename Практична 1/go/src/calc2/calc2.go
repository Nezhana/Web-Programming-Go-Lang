package calc2

import (
	"fmt"
)

type result struct {
	Er [7]element
	Qr float64
}

type element struct {
	Title string
	Value float64
}

func CalculatePart2(hg float64, cg float64, sg float64, ng float64,
	og float64, wp float64, ac float64, qg float64, vc float64) result {
	// Hg := 4.20
	// Cg := 62.1
	// Sg := 3.30
	// Ng := 1.20
	// Og := 6.40
	// Eg := [5]float64{Hg, Cg, Sg, Ng, Og}

	// Wr := 2.00
	// Ac := 0.15
	// Qg := 40.40
	// Vc := 333.3

	Hg := hg
	Cg := cg
	Sg := sg
	Ng := ng
	Og := og
	Eg := [5]float64{Hg, Cg, Sg, Ng, Og}

	Wr := wp
	Ac := ac
	Qg := qg
	Vc := vc

	Kcr := CalcKcr(Wr)
	fmt.Printf("KCR: %7.3f\n", Kcr)
	Ar := CalcAr(Ac, Kcr)
	fmt.Printf("Ar: %7.3f\n", Ar)

	Kgr := CalcKgr(Wr, Ar)
	fmt.Printf("KGR: %7.3f\n", Kgr)
	Er := CalcEr(Wr, Ar, Kgr, Eg)
	for _, er := range Er {
		fmt.Printf("%.2f\t", er)
	}

	Qr := CalcQr2(Qg, Kgr, Wr)
	Vr := CalcVr(Vc, Kcr)

	fmt.Printf("\nQr: %0.3f\tVr: %0.3f", Qr, Vr)

	elsEr := [7]element{{"Hr", Er[0]}, {"Cr", Er[1]}, {"Sr", Er[2]}, {"Nr", Er[3]},
		{"Or", Er[4]}, {"Wr", Er[5]}, {"Ar", Er[6]}}

	calc_result := result{Er: elsEr, Qr: Qr}
	return calc_result
}

func CalcKcr(Wr float64) float64 {
	Kcr := (100.0 - Wr) / 100.0
	return Kcr
}

func CalcAr(Ac float64, Kcr float64) float64 {
	Ar := Ac * Kcr
	return Ar
}

func CalcKgr(Wr float64, Ar float64) float64 {
	Kgr := (100.0 - Wr - Ar) / 100.0
	return Kgr
}

func CalcEr(Wr float64, Ar float64, Kgr float64, Eg [5]float64) [7]float64 {
	var Er [7]float64
	for i := 0; i < 5; i++ {
		Er[i] = Eg[i] * Kgr
	}
	Er[5], Er[6] = Wr, Ar
	return Er
}

func CalcQr2(Qg float64, Kgr float64, Wr float64) float64 {
	Qr := Qg*Kgr - 0.025*Wr
	return Qr
}

func CalcVr(Vc float64, Kcr float64) float64 {
	Vr := Vc * Kcr
	return Vr
}
