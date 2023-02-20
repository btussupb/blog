package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Listening on port localhost:3000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err = tmp.ExecuteTemplate(w, "index", nil); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html", "templates/write.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err = tmp.ExecuteTemplate(w, "write", nil); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
