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

func serverStaticFile(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	log.Printf("filename: %v", name)

	wd, _ := os.Getwd()

	filePath := filepath.Join(wd, "static", name)
	content, _ := os.ReadFile(filePath)

	w.Header().Add("Content-Type", "application/html")
	w.Write(content)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /data", serveJSON)
	mux.HandleFunc("GET /static/{name}", serverStaticFile)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Error starting the server")
	}
}
