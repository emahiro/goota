package main

import (
	"goota/handler"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

const port = "8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Top)
	{
		// qiita
		router.HandleFunc("/goota", handler.GootaIndex).Methods("GET")
		router.HandleFunc("/goota/search", handler.GootaSearchInstances).Methods("GET")

	}

	router.Handle("/", router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("err: %v", err)
	}
}
