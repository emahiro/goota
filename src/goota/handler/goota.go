package handler

import (
	"encoding/json"
	"net/http"

	"goota/service"

	"github.com/labstack/gommon/log"
)

// 検索ページ
func GootaIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Goota"))
}

func GootaSearchInstances(w http.ResponseWriter, r *http.Request) {
	// /goota/search?query=go, golang
	gCli := service.GootaClient{}
	params := r.URL.Query()
	result, err := gCli.GetInstances(&params)
	if err != nil {
		log.Fatalf("GootaGetInstances error. err: %v", err)
	}

	log.Printf("resp: %v", result)
	encoder := json.NewEncoder(w)
	encoder.Encode(result)
}
