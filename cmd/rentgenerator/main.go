package main

import (
	"net/http"
	"path"
	"text/template"

	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/reader"
)

func main() {
	env := environment.LoadRentGeneratorEnvironment()
	rent_data := reader.ReadRentJSON(env)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /gerar", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles(
			path.Join("internal", "template", "base.layout.tmpl"),
			path.Join("internal", "template", "voucher.block.tmpl"),
			path.Join("internal", "template", "bordereau.block.tmpl"),
		)

		tmpl.ExecuteTemplate(w, "layout", rent_data)
	})

	http.ListenAndServe(env.LOCAL_URL, mux)
}
