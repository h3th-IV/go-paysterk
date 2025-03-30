package main

import (
	"log"
	"net/http"

	"github.com/h3th-IV/go-paysterk/paysterk"
)

func main() {
	http.HandleFunc("/go-paysterk", paysterk.WebHookHandler)
	log.Println("webhook server running and listening on :9090")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
