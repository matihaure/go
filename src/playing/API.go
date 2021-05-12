package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login page request received")
		w.Write([]byte("LOGIN PAGE"))
	})

	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start page request received")
		w.Write([]byte("START PAGE"))
	})

	http.ListenAndServe("localhost:3000", mux)

}
