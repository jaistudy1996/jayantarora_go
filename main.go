package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Http handler functions
	http.HandleFunc("/", handleIndex)
	http.Handle("/assets", http.FileServer(http.Dir("static/assets/")))
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/resume", handleResume)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/index.html")
}

func handleFavicon(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/assets/favicon.ico")
}

func handleResume(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/assets/resume.pdf")
}
