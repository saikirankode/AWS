package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//Uploadobject("uploadfile","awsassign","now")
	router := mux.NewRouter()
	router.HandleFunc("/uploadobj", uploadobject).Methods("PUT") // (/uploadobject?filename=&bucketname=&objectname)
	log.Fatal(http.ListenAndServe(":8081", router))

}
