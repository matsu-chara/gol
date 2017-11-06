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
	<h3>current links (<a href="/api/dump">as json</a>)</h3>
    <ul>
      {{ range $key, $value := . }}
      <li>{{ $key }} :<a href="{{ $value }}">{{ $value }}</a></li>
      {{ end }}
	</ul>
	<h3>register new link</h3>
    <form id="register-form">
      <label>key: <input id="register-form-key" name="key" type="text" /></label>
      <label>url: <input id="register-form-value" name="value" type="text" /></label>
      <input type="button" value="register" onclick="doRegister()">
	</form>
	<h3>delete link</h3>
    <form id="delete-form">
      <label>key: <input id="delete-form-key" name="key" type="text" /></label>
      <input type="button" value="delete" onclick="doDelete()">
	</form>
    <script type="text/javascript">
      function doRegister(){
    	let keyInput = document.getElementById("register-form-key");
    	let valueInput = document.getElementById("register-form-value");
		
		var req = new XMLHttpRequest();
		req.onreadystatechange = function() {
		  if (req.readyState == 4) {
			if (req.status == 201) {
			  location.reload();
			} else {	
			  console.error("registration failed. status: " + req.status + ", response:" + req.response);
			  alert("registration failed. status: " + req.status + ", response:" + req.response);
			}
		  }
		};
		req.open("POST", "/" + keyInput.value, true);
		req.setRequestHeader("content-type", "application/x-www-form-urlencoded");
		req.send("value=" + encodeURIComponent(valueInput.value));
    }
	function doDelete(){
    	let keyInput = document.getElementById("delete-form-key");
		
		var req = new XMLHttpRequest();
		req.onreadystatechange = function() {
		  if (req.readyState == 4) {
			if (req.status == 200) {
			  location.reload();
			} else {	
			  console.error("deletion failed. status: " + req.status + ", response:" + req.response);
			  alert("deletion failed. status: " + req.status + ", response:" + req.response);
			}
		  }
		};
		req.open("DELETE", "/" + keyInput.value, true);
		req.send(null);
    }
	</script>
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
