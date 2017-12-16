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
		<style type="text/css">
			table {
				width: 50%;
				border-collapse: collapse;
  		}
			table tr th,	table tr td {
				border: 1px solid #EEE;
			}
			table tr th {
				font-weight: bold;
			}
	  </style>
  </head>
	<body>
		<div class="container">
			<h3>current links (<a href="/api/dump">as json</a>)</h3>
			<table>
				<thead>
					<tr>
						<th>Key</th>
						<th>Link</th>
					</tr>
				</thead>
				<tbody>
					{{ range $key, $value := . }}
					<tr>
						<td>{{ $key }}</td>
						<td><a href="{{ $value.Link }}">{{ $value.Link }}</a></td>
					</tr>
					{{ end }}
				</tbody>
			</table>
			<h3>register new link</h3>
			<form id="register-form">
				<label>key: <input id="register-form-key" name="key" type="text" /></label>
				<label>url: <input id="register-form-link" name="link" type="text" /></label>
				<label>registeredBy(optional): <input id="register-form-registered-by" name="registeredBy" type="text" /></label>
				<label>overwrite: <input id="register-overwrite-is-force" name="is-overwrite" type="checkbox" value="on"/></label>
				<input type="button" value="register" onclick="doRegister()">
			</form>
			<h3>delete link</h3>
			<form id="delete-form">
				<label>key: <input id="delete-form-key" name="key" type="text" /></label>
				<label>registeredBy(optional): <input id="delete-form-registered-by" name="registeredBy" type="text" /></label>
				<input type="button" value="delete" onclick="doDelete()">
			</form>
		</div>
    <script type="text/javascript">
      function doRegister(){
    	let keyInput = document.getElementById("register-form-key");
    	let linkInput = document.getElementById("register-form-link");
		let registeredByInput = document.getElementById("register-form-registered-by");
		let isOverwriteInput = document.getElementById("register-overwrite-is-force");

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
		req.send("link=" + encodeURIComponent(linkInput.value) + "&force=" + (isOverwriteInput.checked).toString() + "&registeredBy=" + encodeURIComponent(registeredByInput.value));
    }
	function doDelete(){
    	let keyInput = document.getElementById("delete-form-key");
			let registeredByInput = document.getElementById("delete-form-registered-by");

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
		req.open("DELETE", "/" + keyInput.value + "?registeredBy=" + encodeURIComponent(registeredByInput.value), true);
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
