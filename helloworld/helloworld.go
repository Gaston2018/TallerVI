package helloworld

import (
	"fmt"
	"io"
	"net/http"
)

func hello() {
	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Holllaaa!!!")
	})
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay una nueva petición")
	io.WriteString(w, "Hola Mundo")
}
