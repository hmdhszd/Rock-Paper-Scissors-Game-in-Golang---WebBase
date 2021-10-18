package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	//-------------------------------------------------------------

	w.Header().Set("Content-Type", "text/html")

	html := `<h1>Hello World inside html tag int the go program!</h1>`
	fmt.Fprint(w, html)

	//-------------------------------------------------------------

	fmt.Fprint(w, "Hello World!")

	//-------------------------------------------------------------

}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080   -->   http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
