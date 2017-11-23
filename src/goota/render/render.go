package render

import (
	"html/template"
	"net/http"
)

func RenderHTML(f string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(f)
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}
