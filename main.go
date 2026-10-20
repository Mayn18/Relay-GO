package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Hello world, the web server
	filePath := filepath.Join(".", "web")
	fmt.Println(filePath)

	fileServer := http.FileServer(http.Dir(filePath))
	log.Println("Server starting on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", fileServer))
}

