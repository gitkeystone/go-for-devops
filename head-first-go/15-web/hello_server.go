package main

import (
	"log"
	"net/http"
)

func write(resp http.ResponseWriter, message string) {
	_, err := resp.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(resp http.ResponseWriter, _ *http.Request) {
	write(resp, "Hello, web!")
}

func frenchHandler(resp http.ResponseWriter, _ *http.Request) {
	write(resp, "Salut, web!")
}

func hindiHandler(resp http.ResponseWriter, _ *http.Request) {
	write(resp, "Namaste, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
