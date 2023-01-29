package main

import (
	"log"
	"net/http"
	"text/template"
)

// Temps is template structs.
type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body>NO TEMPALTES.</body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

// setup template function.
func setupTemp() *Temps {
	temps := new(Temps)
	temps.notemp = notemp()

	// set index template.
	indx, er := template.ParseFiles("templates/index.html")
	if er != nil {
		indx = temps.notemp
	}

	temps.indx = indx

	// set hello template.
	helo, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		helo = temps.notemp
	}
	temps.helo = helo

	return temps
}

// index handler
func index(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler
func hello(w http.ResponseWriter, req *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	temps := setupTemp()
	// index handling
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		index(w, req, temps.indx)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		hello(w, req, temps.helo)
	})
	http.ListenAndServe("", nil)
}