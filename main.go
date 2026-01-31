package main

import (
	"net/http"
	"log"
	"github.com/Wahbi8/PM_Golang/apis"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /email/invoice", apis.SendEmailApi)

	log.Fatal(http.ListenAndServe(":1212", mux))

}