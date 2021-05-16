package controller

import (
	"net/http"
)

func Register() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/dbtest", dbtest)
	mux.HandleFunc("/enginedbtest", enginedbtest)
	mux.HandleFunc("/", crud())

	return mux

}
