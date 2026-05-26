package service

import (
	"html"
	"net/http"
	"text/template"

	"github.com/NYTimes/gizmo/examples/nyt"
)

func (s *SimpleService) GetCats(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

var tempFuncs = template.FuncMap{"unescape": html.UnescapeString}
var catsTemplate = template.Must(template.New("dash").Funcs(tempFuncs).Parse(dashHTML))

type catList struct {
	Articles []*nyt.SemanticConceptArticle
}

const dashHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>NYT Cat Articles</title>
  </head>
  <body>
	<h1>Recent nytimes.com Articles about "Cats"</h1>
	<ol>
		{{ range .Articles }}
			<li class="article">
				<h2>{{ unescape .Title }}</h2>
				<h5>{{ unescape .Byline }}</h5>
				<p><a target="none" href="{{ .Url }}">{{ .Url }}</a></p>
				<p>{{ unescape .Body }}</p>
			</li>
		{{ end }}
	</ol>
  </body>
</html>`
