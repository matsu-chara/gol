package usecase

import (
	"html/template"
	"net/http"

	"github.com/matsu-chara/gol/operations"
)

var dumpTemplate = template.Must(template.New("gol").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>gol</title>
  </head>
  <body>
  <ul>
   {{ range $key, $value := . }}
   <li>{{ $key }} :<a href="{{ $value }}">{{ $value }}</a></li>
   {{ end }}
  </ul>
  </body>
</html>
`))

// DumpAsHTML dumps all links in kvs as html
func DumpAsHTML(filepath string, w http.ResponseWriter) {
	dumped, err := operations.RunDump(filepath)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = dumpTemplate.ExecuteTemplate(w, "gol", dumped)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	return
}
