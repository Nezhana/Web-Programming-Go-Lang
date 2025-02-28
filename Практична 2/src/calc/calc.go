package calc

import (
	"math"
)

type Result struct {
	OptionTitle string
	K           float64
	E           float64
}

func CalculatePart1(option int, wp float64, ap float64,
	Q float64, B float64, n float64) Result {

	// Для вугілля (робоча маса палива)
	//  Hr,    Cr,    Sr,   Nr,   Or,   Wr,    Ar    Vr
	// 3.50, 52.49, 2.85, 0.97, 4.99, 10.00, 25.20, 25.92

	// Для мазута (горюча маса палива)
	//  Hg,    Cg,    Sg,   Ng,   Og,   Wr,   As     V
	// 11.20, 85.50, 2.50, 0.80, 0.80, 2.00, 0.15, 333.3

	// Для природного газу
	//  CH4,  C2H6,  C3H8, C4H10, CO2,  N2    ro     X
	// 98.90, 0.12, 0.011, 0.01, 0.06, 0.90, 0.723, 0.0

	Wp := wp
	Ap := ap

	var a, kS, Gvin, Qr, Ar, Krc float64
	var optionTitle string

	// option := 2

	if option == 2 {
		a = 1.0
		kS = 0.0
		Krc = CalcKrc(Wp)
		Ar = Ap * Krc
		Qr = Q * ((100-Wp-Ar)/100 - 0.025*Wp)
		Gvin = 100 / (100 - Wp - Ar)
		optionTitle = "Мазут"
	} else if option == 3 {
		a = 0.0
		kS = 0.0
		Gvin = 0.0
		Ar = 0.0
		Qr = Q
		optionTitle = "Природний газ"
	} else {
		a = 0.80
		kS = 0.0
		Gvin = 1.5
		Qr = Q
		Ar = Ap
		optionTitle = "Вугілля"
	}

	k := CalcK(Qr, a, Ar, Gvin, n, kS)
	E := CalcE(Qr, k, B)

	calc_result := Result{OptionTitle: optionTitle, K: k, E: E}
	return calc_result
}

func CalcKrc(Wp float64) float64 {
	Krc := (100.0 - Wp) / 100.0
	return Krc
}

func CalcKrg(Wp float64, Ap float64) float64 {
	Krg := 100.0 / (100.0 - Wp - Ap)
	return Krg
}

func CalcKcg(Ac float64) float64 {
	Krg := 100.0 / (100.0 - Ac)
	return Krg
}

func CalcK(Qr float64, a float64, Ap float64, Gvin float64, n float64, kS float64) float64 {
	k := (math.Pow(10.0, 6.0)/Qr)*a*(Ap/(100-Gvin))*(1-n) + kS
	return k
}

func CalcE(Qr float64, k float64, B float64) float64 {
	E := math.Pow(10.0, -6.0) * k * Qr * B
	return E
}

func CalcQcr(Wp float64) float64 {
	Qcr := (100.0 - Wp) / 100.0
	return Qcr
}

func CalcQgr(Wp float64, Ap float64) float64 {
	Qgr := (100.0 - Wp - Ap) / 100.0
	return Qgr
}
