package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Holllaaa!!!")
	})
	http.HandleFunc("/", handler)
	http.ListenAndServe("http://ec2-18-219-147-32.us-east-2.compute.amazonaws.com/", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay una nueva petición")
	io.WriteString(w, "Hola Mundo")
}
