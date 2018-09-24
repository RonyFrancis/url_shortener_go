package main

import (
	"net/http"

	"github.com/RonyFrancis/url_shortener_go/controller"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/shortenUrl", shortener.AddUrlHandler)
	myRouter.Methods("POST")
	http.ListenAndServe(":8001", myRouter)
}
