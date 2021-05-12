package controller

import (
	"encoding/json"
	"fmt"
	"github.com/matihaure/model"
	"github.com/matihaure/views"
	"net/http"
)


func create(w http.ResponseWriter, r *http.Request){
		data := views.PostRequest{}
		json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
		if err:= model.CreateTodo(data.Name,data.Todo); err!=nil{
			w.Write([]byte("Some error"))
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
}

func readAll(w http.ResponseWriter){

	records,err := model.ReadAll()

	if err!= nil{
		w.Write([]byte("Some error"))
		return
	}
		json.NewEncoder(w).Encode(records)


	w.WriteHeader(http.StatusOK)


}

func readbyid(w http.ResponseWriter, data string){

	records,err := model.ReadbyId(data)

	if err!= nil{
		w.Write([]byte("Some error"))
		return
	}
	json.NewEncoder(w).Encode(records)


	w.WriteHeader(http.StatusOK)


}

func deletebyid(w http.ResponseWriter, data string){

	err:= model.DeleteById(data)

	if err!= nil{
		w.Write([]byte("Some error"))
		return
	}


	w.WriteHeader(http.StatusOK)


}

func crud() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){

		if r.Method == http.MethodPost{
		create(w,r)
		}

		if r.Method == http.MethodGet{

			data:= r.URL.Query().Get("name")

			if data ==""{
				readAll(w)
			}
			readbyid(w,data)
		}

		if r.Method == http.MethodDelete{
			data:= r.URL.Query().Get("name")
			deletebyid(w,data)
		}

	}

}
