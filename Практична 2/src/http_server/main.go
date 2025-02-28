package main

import (
	"calc2/src/calc"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type elements struct {
	Option int
	El     [8]string
}

type htmlData struct {
	Data       calc.Result
	IsAnswered bool
	ElTitles   elements
}

func check(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

var OPTION int

func save_option(option int) {
	OPTION = option
}

func get_option() int {
	return OPTION
}

func getElTitles(option int) elements {

	coal_el := elements{1, [8]string{"Hр", "Cр", "Sр", "Nр",
		"Oр", "Wр", "Aр", "Vр"}}
	oil_el := elements{2, [8]string{"Hг", "Cг", "Sг", "Nг",
		"Oг", "Wр", "Aг", "V"}}
	gas_el := elements{3, [8]string{"CH4", "C2H6", "C3H8",
		"C4H10", "CO2", "N2", "ρ", "X"}}

	var elTitles elements
	if option == 2 {
		elTitles = oil_el
	} else if option == 3 {
		elTitles = gas_el
	} else {
		elTitles = coal_el
	}

	return elTitles
}

func calculateHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Error parsing form", http.StatusBadRequest)
	}

	// option := 1
	// option, _ := strconv.Atoi(request.FormValue("option")) // Get selected option
	option := get_option()
	fmt.Printf("\noption: %d", option)
	wp, _ := strconv.ParseFloat(request.FormValue("Wp"), 64)
	ap, _ := strconv.ParseFloat(request.FormValue("Ap"), 64)
	Q, _ := strconv.ParseFloat(request.FormValue("Q"), 64)
	B, _ := strconv.ParseFloat(request.FormValue("B"), 64)
	n, _ := strconv.ParseFloat(request.FormValue("n"), 64)

	result := calc.CalculatePart1(option, wp, ap, Q, B, n)

	fmt.Printf("\nk: %0.3f\tE: %0.3f\n", result.K, result.E)

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(writer, "Error loading template", http.StatusInternalServerError)
		return
	}

	elTitles := getElTitles(option)

	calcData := htmlData{Data: result, IsAnswered: true, ElTitles: elTitles}

	// Send data to template
	err = tmpl.Execute(writer, calcData)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
	}

}

func viewHandler(writer http.ResponseWriter, request *http.Request) {

	optionStr := request.URL.Query().Get("option") // Get the fuel type option
	option, err := strconv.Atoi(optionStr)
	if err != nil {
		option = 1 // Default to Вугілля if not specified
	}

	save_option(option)

	data1 := calc.CalculatePart1(option, 0.0, 0.0, 0.0, 0.0, 0.0)

	elTitles := getElTitles(option)

	startData := htmlData{Data: data1, IsAnswered: false, ElTitles: elTitles}
	fmt.Printf("\nk: %0.3f\tE: %0.3f\n", data1.K, data1.E)

	html, err := template.ParseFiles("index.html")
	check(err)
	err = html.Execute(writer, startData)
	check(err)
}

func startPageHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("start-page.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

// Serve static files with correct MIME types
func staticFileServer(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/static/")
	fullPath := filepath.Join("static", path)

	http.ServeFile(w, r, fullPath)
}

func main() {
	http.HandleFunc("/static/", staticFileServer) // Fix MIME type issue
	http.HandleFunc("/startpage", startPageHandler)
	http.HandleFunc("/calculator", viewHandler)
	http.HandleFunc("/calculate", calculateHandler)

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
