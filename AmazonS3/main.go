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
	router.HandleFunc("/listobj", Listobject).Methods("GET")     //(/listobj?bucketname=)
	router.HandleFunc("/downloadobj", Downloadobject).Methods("GET") //(/downloadobj?bucketname=&objectname=)
	router.HandleFunc("/deleteobj", Deleteobject).Methods("DELETE")  //(/deleteobj?bucketname=&objectname=)
	log.Fatal(http.ListenAndServe(":8081", router))

}
