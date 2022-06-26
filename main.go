package main

import (
		"net/http"
		"log"
		"github.com/gorilla/mux"
)


func main()  {
	router := mux.NewRouter()
	InitialMigration()
	router.HandleFunc("/",Home).Methods("POST")
	router.HandleFunc("/splitbill",CreateTranscation).Methods("POST")
	router.HandleFunc("/showbill/{number}/{intent}",ShowTranscation).Methods("GET")
	router.HandleFunc("/pay/{number}/{id}",PayTranscation).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10021", router))
}