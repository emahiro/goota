package render

import (
	"html/template"
	"net/http"
)

func RenderHTML(filename string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	if err := t.ExecuteTemplate(w, filename, data); err != nil {
		return err
	}
	return nil
}
