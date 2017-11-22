package service

import (
	"encoding/json"
	"goota/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/gommon/log"
)

var (
	QIITA_API_PATH  = "https://qiita.com/api/v2/items"
	PAGE            = "1"
	MAX_PER_PAGE    = "20"
	MIN_STOCK_COUNT = 100
)

type GootaService interface {
	GetInstances(r *http.Request) (*http.Response, error)
}

type GootaClient struct {
	GootaService
}

func (g GootaClient) new() http.Client {
	return http.Client{}
}

func (g GootaClient) GetInstances(v *url.Values) (*[]model.Goota, error) {
	var query string
	rawTags := v.Get("query")
	query = buildQuery(rawTags)

	values := url.Values{}
	values.Add("page", PAGE)
	values.Add("per_page", MAX_PER_PAGE)
	values.Add("query", query)

	cli := g.new()
	resp, err := cli.Get(QIITA_API_PATH + "?" + values.Encode())
	if err != nil {
		log.Fatalf("request error. err: %v", err)
		return nil, err
	}

	body := resp.Body
	defer body.Close()
	b, _ := ioutil.ReadAll(body)
	var result []model.Goota
	if err := json.Unmarshal(b, &result); err != nil {
		log.Fatalf("json unmarshal error. err: %v", err)
	}

	return &result, nil
}

func buildQuery(rawTags string) string {
	var q string
	tags := extractTags(rawTags)
	l := len(tags)
	for i, t := range tags {
		if i != l-1 {
			q += "tag:" + t + " OR "
			continue
		}
		q += "tag:" + t
		break
	}
	return q
}

func extractTags(t string) []string {
	return strings.Split(t, ",")
}
