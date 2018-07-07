package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start HTTP server: ", err)
	}
}
