package main

import (
	"net/http"

	"github.com/RonyFrancis/url_shortener_go/controller"
	"github.com/RonyFrancis/url_shortener_go/model"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	myRouter := mux.NewRouter()
	model.InitalMigration()
	myRouter.HandleFunc("/shortenUrl", shortener.AddUrlHandler)
	myRouter.Methods("POST")
	http.ListenAndServe(":8001", myRouter)
}
