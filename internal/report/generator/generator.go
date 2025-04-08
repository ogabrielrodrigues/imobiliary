package generator

import (
	"net/http"
	"path"
	"text/template"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared"
)

func Generate(url string, houses []kind.House) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /report", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles(path.Join("template", "report.template.tmpl"))

		data := kind.Tmpl{
			Date:   shared.CurrentDate(),
			Houses: houses,
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(url, mux)
}
