package main

import (
	"calc/src/calc1"
	"calc/src/calc2"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func executeTemplate(text string, data interface{}) {
// 	tmpl, err := template.New("tmpl").Parse(text)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, data)
// 	check(err)
// }

func calculateHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Error parsing form", http.StatusBadRequest)
	}

	hp, _ := strconv.ParseFloat(request.FormValue("Hp"), 64)
	cp, _ := strconv.ParseFloat(request.FormValue("Cp"), 64)
	sp, _ := strconv.ParseFloat(request.FormValue("Sp"), 64)
	np, _ := strconv.ParseFloat(request.FormValue("Np"), 64)
	op, _ := strconv.ParseFloat(request.FormValue("Op"), 64)
	wp, _ := strconv.ParseFloat(request.FormValue("Wp"), 64)
	ap, _ := strconv.ParseFloat(request.FormValue("Ap"), 64)

	results := calc1.CalculatePart1(hp, cp, sp, np, op, wp, ap)

	fmt.Println("\nEc:")
	for _, ec := range results.Ec {
		fmt.Printf("%.2f\t", ec)
	}
	fmt.Printf("\nQr: %0.3f\tQc: %0.3f\tQg: %0.3f\n", results.Qr, results.Qc, results.Qg)

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(writer, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Send data to template
	err = tmpl.Execute(writer, results)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
	}
}

func calculateHandler2(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Error parsing form", http.StatusBadRequest)
	}

	hg, _ := strconv.ParseFloat(request.FormValue("Hg"), 64)
	cg, _ := strconv.ParseFloat(request.FormValue("Cg"), 64)
	sg, _ := strconv.ParseFloat(request.FormValue("Sg"), 64)
	ng, _ := strconv.ParseFloat(request.FormValue("Ng"), 64)
	og, _ := strconv.ParseFloat(request.FormValue("Og"), 64)
	wp, _ := strconv.ParseFloat(request.FormValue("Wp"), 64)
	ac, _ := strconv.ParseFloat(request.FormValue("Ac"), 64)
	qg, _ := strconv.ParseFloat(request.FormValue("Qg"), 64)
	vc, _ := strconv.ParseFloat(request.FormValue("Vc"), 64)

	results := calc2.CalculatePart2(hg, cg, sg, ng, og, wp, ac, qg, vc)

	fmt.Println("\nEr:")
	for _, er := range results.Er {
		fmt.Printf("%.2f\t", er)
	}
	fmt.Printf("\nQr: %0.3f\n", results.Qr)

	tmpl, err := template.ParseFiles("index2.html")
	if err != nil {
		http.Error(writer, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Send data to template
	err = tmpl.Execute(writer, results)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {

	data1 := calc1.CalculatePart1(0, 0, 0, 0, 0, 0, 0)
	fmt.Printf("\nQr: %0.3f\tQc: %0.3f\tQg: %0.3f\n", data1.Qr, data1.Qc, data1.Qg)

	html, err := template.ParseFiles("index.html")
	check(err)
	err = html.Execute(writer, data1)
	check(err)
}

func viewHandler2(writer http.ResponseWriter, request *http.Request) {

	data2 := calc2.CalculatePart2(0, 0, 0, 0, 0, 0, 0, 0, 0)
	fmt.Printf("\nQr: %0.3f\n", data2.Qr)

	html, err := template.ParseFiles("index2.html")
	check(err)
	err = html.Execute(writer, data2)
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
	http.HandleFunc("/calculator", viewHandler)
	http.HandleFunc("/calculator2", viewHandler2)
	http.HandleFunc("/calculate", calculateHandler)   // Handle form submission
	http.HandleFunc("/calculate2", calculateHandler2) // Handle form submission

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
