package main

import (
		"net/http"
		"log"
		"github.com/gorilla/mux"
		"view"
		"models"
)


func main()  {
	router := mux.NewRouter()
	models.InitialMigration()
	router.HandleFunc("/",view.Home).Methods("POST")
	router.HandleFunc("/splitbill",view.CreateTranscation).Methods("POST")
	router.HandleFunc("/showbill/{number}/{intent}",view.ShowTranscation).Methods("GET")
	router.HandleFunc("/pay/{number}/{id}",view.PayTranscation).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10021", router))
}