package main

import (
	"github.com/matihaure/controller"
	"github.com/matihaure/model"
	"net/http"
)

func main() {

	mux := controller.Register()
	model.Connect()

	http.ListenAndServe(":3000", mux)



	//result, err := db.Query("show databases")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for result.Next()

}
