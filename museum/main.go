package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"danieletortora.com/go/museum/data"
)

type DataResponse struct {
	Message string `json:"message"`
}

func serveJSON(w http.ResponseWriter, r *http.Request) {

	data := DataResponse{Message: "Hello from handler"}

	encoded, _ := json.Marshal(data)

	w.Write(encoded)
}

func createFileServer() http.Handler {
	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "static")

	fs := http.FileServer(http.Dir(filePath))

	return fs
}

func serveGoTemplate(w http.ResponseWriter, r *http.Request) {
	data := data.GetAll()

	t, _ := template.ParseFiles("templates/index.tmpl")

	t.Execute(w, data)

}

func main() {
	mux := http.NewServeMux()

	// Server JSON responses
	mux.HandleFunc("GET /data", serveJSON)
	// Serving Go Template
	mux.HandleFunc("GET /template", serveGoTemplate)
	// Server Static files
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", createFileServer())
	})

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Error starting the server")
	}
}
