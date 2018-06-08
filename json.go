package main

import (
	"net/http"
	//"encoding/json"
)

type curso struct {
	Title         string
	NumerosVideos int
}

func main() {
	//http.HandleFunc("/",func(){
	//
	//})

	http.ListenAndServe(":8080", nil)
}
