package controller

import (
	"encoding/json"
	"fmt"
	"github.com/matihaure/views"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request){

		if r.Method == http.MethodGet {
			fmt.Println("Recibi un GET")

			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)

		}

}
