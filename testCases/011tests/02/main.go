package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RootEndpoint( responce http.ResponseWriter,req *http.Request){
	responce.WriteHeader(200)
	responce.Write([]byte("Hello world"))

}


func main (){
	router := mux.NewRouter()
	router.HandleFunc("/",RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345",router))
}
