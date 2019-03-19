package usecase

import (
	"html/template"
	"net/http"

	"github.com/matsu-chara/gol/operations"
)

var dumpTemplate = template.Must(template.ParseFiles("templates/dump_as_html.html"))

// DumpAsHTML dumps all links in kvs as html
func DumpAsHTML(filepath string, w http.ResponseWriter) {
	dumped, err := operations.RunDump(filepath)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = dumpTemplate.ExecuteTemplate(w, "dump_as_html.html", dumped)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	return
}
