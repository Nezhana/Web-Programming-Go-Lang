package calc1

type element struct {
	Title string
	Value float64
}

type result struct {
	Ec [6]element
	Eg [5]element
	Qr float64
	Qc float64
	Qg float64
}

func CalculatePart1(hp float64, cp float64, sp float64, np float64, op float64, wp float64, ap float64) result {
	// Hp := 4.20
	// Cp := 62.1
	// Sp := 3.30
	// Np := 1.20
	// Op := 6.40
	// Wp := 7.00
	// Ap := 15.80

	Hp := hp
	Cp := cp
	Sp := sp
	Np := np
	Op := op
	Wp := wp
	Ap := ap

	Krc := CalcKrc(Wp)
	// fmt.Printf("KRC: %7.3f\n", Krc)

	Krg := CalcKrg(Wp, Ap)
	// fmt.Printf("KRG: %7.3f\n", Krg)

	// fmt.Println(reflect.TypeOf(Hp))

	// fmt.Println("Er:")
	Er := [7]float64{Hp, Cp, Sp, Np, Op, Wp, Ap}
	// for _, er := range Er {
	// 	fmt.Printf("%.2f\t", er)
	// }

	// fmt.Println("\nEc:")
	Ec := CalcEc(Krc, Er)
	// for _, ec := range Ec {
	// 	fmt.Printf("%.2f\t", ec)
	// }

	// fmt.Println("\nEg:")
	Eg := CalcEg(Krg, Er)
	// for _, eg := range Eg {
	// 	fmt.Printf("%.2f\t", eg)
	// }

	Kn := [4]float64{339.0, 1030.0, 108.8, 25.0}

	Qr := CalcQr(Er, Kn)
	Qc, Qg := CalcQc(Qr, Krc, Er), CalcQg(Qr, Krg, Er)

	// fmt.Printf("\nQr: %0.3f\tQc: %0.3f\tQg: %0.3f", Qr, Qc, Qg)

	elsEc := [6]element{{"Hc", Ec[0]}, {"Cc", Ec[1]}, {"Sc", Ec[2]},
		{"Nc", Ec[3]}, {"Oc", Ec[4]}, {"Ac", Ec[5]}}

	elsEg := [5]element{{"Hg", Eg[0]}, {"Cg", Eg[1]}, {"Sg", Eg[2]},
		{"Ng", Eg[3]}, {"Og", Ec[4]}}

	// for _, eg := range elsEg {
	// 	fmt.Printf("%s : %.2f\t",eg.title, eg.value)
	// }

	calc_result := result{Ec: elsEc, Eg: elsEg, Qr: Qr, Qc: Qc, Qg: Qg}

	return calc_result
}

// ------------------------------------------------

// strings to numbers => page 81
// random num int generation => page 91
// loops => page 95
// formating => default: fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
/*
  fmt.Printf("A float: %f\n", 3.1415)
  fmt.Printf("An integer: %d\n", 15)
  fmt.Printf("A string: %s\n", "hello")
  fmt.Printf("A boolean: %t\n", false)
  fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
  fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
  fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
  fmt.Printf("Percent sign: %%\n")

  fmt.Printf("%12s | %2d\n", "Stamps", 50)

  fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
*/
//funcs => page 125
//arrays => page 187

// ------------------------------------------------

func CalcKrc(Wp float64) float64 {
	Krc := 100.0 / (100.0 - Wp)
	return Krc
}

func CalcKrg(Wp float64, Ap float64) float64 {
	Krg := 100.0 / (100.0 - Wp - Ap)
	return Krg
}

func CalcEc(Krc float64, Er [7]float64) [6]float64 {
	var Ec [6]float64
	for i := 0; i < 7; i++ {
		Ec[i] = Er[i] * Krc
		if i == 5 {
			Ec[i] = Er[i+1] * Krc
			break
		}
	}
	return Ec
}

func CalcEg(Krg float64, Er [7]float64) [5]float64 {
	var Eg [5]float64
	for i := 0; i < 6; i++ {
		Eg[i] = Er[i] * Krg
		if i == 4 {
			break
		}
	}
	return Eg
}

func CalcQr(Er [7]float64, Kn [4]float64) float64 {
	Qr := (Kn[0] * Er[1]) + (Kn[1] * Er[0]) - Kn[2]*(Er[4]-Er[2]) - (Kn[3] * Er[5])
	return Qr
}

func CalcQc(Qr float64, Krc float64, Er [7]float64) float64 {
	Qc := (Qr + 0.025*Er[5]) * Krc
	return Qc
}

func CalcQg(Qr float64, Krg float64, Er [7]float64) float64 {
	Qg := (Qr + 0.025*Er[5]) * Krg
	return Qg
}
