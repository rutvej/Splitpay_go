package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    gateway_information := string(body)

	fmt.Println(gateway_information)
    fmt.Fprintf(w,"hello")

}


func CreateTranscation(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    payload := string(body)
	ids := Save(payload)
	fmt.Println(ids)
    fmt.Fprintf(w,ids)
}

func ShowTranscation(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	payload := GetAll(params["number"],params["intent"])
	fmt.Println(payload)
    fmt.Fprintf(w,payload)
}
func PayTranscation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	go UpdateStatus(params["id"],params["number"])
	fmt.Println("success")
    fmt.Fprintf(w,`{"status":"success"}`)
}



