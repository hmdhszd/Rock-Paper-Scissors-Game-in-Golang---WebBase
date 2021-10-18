package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	log.Println("Starting web server on port 8080   -->   http://127.0.0.1:8080")

	http.ListenAndServe(":8080", nil)

}

func renderTemplate(w http.ResponseWriter, page string) {

	//Parse the html file
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	//Execute the html file
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
