package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"log"

	"github.com/gorilla/mux"
	//"github.com/youthtrouble/Interest-Calculator/calculator"
)

var (
	//Fs serves static files on the server
	fs = http.FileServer(http.Dir("./templates/css/"))
	js = http.FileServer(http.Dir("./templates/js/"))
	//Fimg servers image files on the server
	//Fimg serves images to the server
	fimg = http.FileServer(http.Dir("./templates/img/"))

	tpl   *template.Template
	tmpl  = template.New("")
	port string
)

func init() {
	// Get serve port from environment
	// Allows for easy port change
	port = os.Getenv("EMS_PORT")
	//Default to 9000
	if port == "" {
		port = "9000"
	}
	// Parses gohtml files in templates directory
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

//Savecalc func
func savecalc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		savecalcGet(w, r)
	// case "POST":
	// 	savecalcPost(w, r)

	}
}

func savecalcGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println("error loading template", err)
		w.WriteHeader(http.StatusNotFound)
	}
}


func main() {
	// Register routers
	router := mux.NewRouter()
	router.HandleFunc("/", savecalc).Methods("GET", "POST")
	//localhost can be omitted
	fmt.Printf("Serving on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
