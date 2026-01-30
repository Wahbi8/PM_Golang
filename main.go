package main

import (
	"net/http"
	"log"
	"github.com/Wahbi8/PM_Golang/apis"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/url", apis.SendEmailApi)

	log.Fatal(http.ListenAndServe(":1212", mux))

}