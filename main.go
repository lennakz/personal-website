package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*/*"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "main.html", "")
	if err != nil {
		log.Fatal("Cannot get view ", err)
	}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", indexHandler)
	fmt.Println("Server running on port 8081 ...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
