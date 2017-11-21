package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"goota/model"
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
	resp, err := gCli.GetInstances(&params)
	if err != nil {
		log.Fatalf("GootaGetInstances error. err: %v", err)
	}

	body := resp.Body
	defer body.Close()
	b, _ := ioutil.ReadAll(body)
	var q []model.Goota
	if err := json.Unmarshal(b, &q); err != nil {
		log.Fatalf("json unmarshal error. err: %v", err)
	}

	log.Printf("resp: %v", q)
}
