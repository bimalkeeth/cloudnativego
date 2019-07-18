package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",indexFunction)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal("Error ocuured")
	}
}

func indexFunction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_,err:=fmt.Fprintf(w,"Hello cloud native go")
	if err!=nil{
		log.Fatal("error in response")
	}
}

