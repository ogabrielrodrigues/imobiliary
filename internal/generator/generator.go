package generator

import (
	"net/http"
	"path"
	"text/template"

	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/util"
)

func Generate(houses []kind.House) {
	env := environment.LoadEnvironment()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /report", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles(path.Join("tmpl", "tmpl.html"))

		data := kind.Tmpl{
			Date:   util.CurrentDate(),
			Houses: houses,
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(env.LOCAL_URL, mux)
}
