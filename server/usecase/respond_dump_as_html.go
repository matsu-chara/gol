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
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0">
  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.100.2/css/materialize.min.css">
  <!-- https://favicon.io/emoji-favicons/paperclip -->
  <link rel="shortcut icon" href="data:image/png;base64,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"
  />
</head>

<body>
  <nav class="light-blue darken-1" role="navigation">
    <div class="nav-wrapper container">
      <span class="brand-logo grey-text text-lighten-5"><i class="material-icons">link</i>gol - private URL shortner</span>
    </div>
  </nav>
  <div class="container">
    <h4><i class="material-icons">view_list</i> current links (<a href="/api/dump">as json</a>)</h4>
    <table class="responsive-table">
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
    <div class="row">
      <h4>register new link</h4>
      <form id="register-form" class="col s12">
        <div class="row">
          <div class="input-field col s4">
            <input id="register-form-key" type="text" placeholder="short_name">
            <label for="register-form-key">key</label>
          </div>
          <div class="input-field col s4">
            <input id="register-form-link" type="text" placeholder="http://some.long.url" />
            <label for="register-form-link">url</label>
          </div>
          <div class="input-field col s4">
            <input id="register-form-registered-by" type="text" placeholder="your_name" class="tooltipped" data-position="bottom" data-delay="50" data-tooltip="If you specify this, you can lock it with your name" />
            <label for="register-form-registered-by">registeredBy (optional)</label>
          </div>
        </div>
        <p>
          <input id="register-overwrite-is-force" type="checkbox" class="filled-in" value="on" />
          <label for="register-overwrite-is-force">overwrite existing link</label>
        </p>
        <button class="btn waves-effect waves-light" type="button" value="register" onclick="doRegister()">Create Link</button>
      </form>
    </div>
    <div class="row">
      <h4>delete link</h4>
      <form id="delete-form" class="col s12">
        <div class="row">
          <div class="input-field col s4">
            <input id="delete-form-key" type="text" />
            <label for="delete-form-key">key</label>
          </div>
          <div class="input-field col s4">
            <input id="delete-form-registered-by" name="registeredBy" type="text" class="tooltipped" data-position="bottom" data-delay="50" data-tooltip="The name you specified when you registered. Deletion will fail if a different name is specified" />
            <label for="delete-form-registered-by">registeredBy</label>
          </div>
        </div>
        <button class="btn waves-effect waves-light" type="button" value="delete" onclick="doDelete()">Delete Link</button>
      </form>
    </div>
  </div>
  <script type="text/javascript" src="https://code.jquery.com/jquery-3.2.1.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.100.2/js/materialize.min.js"></script>
  <script type="text/javascript">
    function doRegister() {
      let keyInput = document.getElementById("register-form-key");
      let linkInput = document.getElementById("register-form-link");
      let registeredByInput = document.getElementById("register-form-registered-by");
      let isOverwriteInput = document.getElementById("register-overwrite-is-force");

      var req = new XMLHttpRequest();
      req.onreadystatechange = function() {
        if (req.readyState == 4) {
          if (req.status == 201) {
            window.scrollTo(0, 0);
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

    function doDelete() {
      let keyInput = document.getElementById("delete-form-key");
      let registeredByInput = document.getElementById("delete-form-registered-by");

      var req = new XMLHttpRequest();
      req.onreadystatechange = function() {
        if (req.readyState == 4) {
          if (req.status == 200) {
            window.scrollTo(0, 0);
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
