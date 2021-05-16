package main

import (
	"net/http"

	"github.com/matihaure/controller"
	"github.com/matihaure/model"
)

func main() {

	mux := controller.Register()
	model.Connect()

	http.ListenAndServe(":3000", mux)

}
