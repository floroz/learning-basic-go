package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type DataResponse struct {
	Message string `json:"message"`
}

func serveJSON(w http.ResponseWriter, r *http.Request) {

	data := DataResponse{Message: "Hello from handler"}

	encoded, _ := json.Marshal(data)

	w.Write(encoded)
}

func main() {
	mux := http.NewServeMux()

	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "static")

	fs := http.FileServer(http.Dir(filePath))

	// Server JSON responses
	mux.HandleFunc("GET /data", serveJSON)
	// Server Static files
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Error starting the server")
	}
}
