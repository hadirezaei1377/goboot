package html2

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template.html"))

	data := struct {
		Title   string
		Heading string
		Message string
	}{
		Title:   "Welcome to My Website",
		Heading: "Hello, World!",
		Message: "This is a dynamic website powered by Golang.",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
