package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			fmt.Println("Recibi un GET")

			data := Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)

		}

	})

	http.ListenAndServe(":3000", mux)

}
