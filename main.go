package main

import (
		"net/http"
		"log"
		"github.com/gorilla/mux"
		"os"
)


func main()  {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	InitialMigration()
	router.HandleFunc("/",Home).Methods("GET")
	router.HandleFunc("/splitbill",CreateTranscation).Methods("POST")
	router.HandleFunc("/showbill/{number}/{intent}",ShowTranscation).Methods("GET")
	router.HandleFunc("/pay/{number}/{id}",PayTranscation).Methods("PUT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}