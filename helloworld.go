package main

import(
    "net/http"
    "io"
)

func main()
{
    http.ListenAndServer("http://ec2-18-219-147-32.us-east-2.compute.amazonaws.com/",nil)
}

func handler(w http.ResponseWriter, r *http.Request){

    io.WriteString(w, "Hola Mundo")
}