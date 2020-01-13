package main

import (
	"io"
	"log"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.Proto, req.Header, req.Body)
	_, _ = io.WriteString(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/", headers)
	log.Fatal(http.ListenAndServe(":30080", nil))
}
