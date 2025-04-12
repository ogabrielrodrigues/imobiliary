package service

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/internal/reader"
)

func RentMaker() {
	rent_data := reader.ReadRentJSON(&kind.RentGeneratorEnvironment{
		RENT_PATH: "./data/alugueis.json",
	})

	tmpl, _ := template.ParseFiles(
		path.Join("internal", "template", "base.layout.tmpl"),
		path.Join("internal", "template", "voucher.block.tmpl"),
		path.Join("internal", "template", "bordereau.block.tmpl"),
	)

	f, err := os.Create("./output.pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "layout", rent_data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Template executed successfully")
}
