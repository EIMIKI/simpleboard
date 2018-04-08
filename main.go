package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/simple", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.Form.Get("name")
		text := r.Form.Get("text")
		if name != "" && text != "" {
			toDb(name, text)
		}

		data := fromDb()
		tmpl, _ := template.ParseFiles(filepath.Join("templates", "board.html"))
		tmpl.Execute(w, data)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
