package main

import (
	// "fmt"
	"log"
	"net/http"
)


func main(){
	port := 9090
	// sm := http.NewServeMux()
	// http.ha
	log.Print("fsfssfs")
	log.Printf("Listen on port %v" ,port)
	http.ListenAndServe(":9090",nil)

}