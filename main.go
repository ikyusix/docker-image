package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)

	fmt.Println("Server Start at localhost with port 9000")
	http.ListenAndServe(":9000", nil)
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Only Accept Method GET", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Hello World"))
}
