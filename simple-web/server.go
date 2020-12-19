package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Print("failed to write")
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
