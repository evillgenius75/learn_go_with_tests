package main

import (
	"di"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(di.MyGreeterHandler)))
}
