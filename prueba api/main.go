package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?id=3433955&APPID=af4c25b26d16e7216792d5c73ffc6584")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
