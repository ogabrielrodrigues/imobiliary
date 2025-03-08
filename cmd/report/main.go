package main

import (
	"fmt"
	"path"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/generator"
	"github.com/ogabrielrodrigues/imobiliary/internal/reader"
	"github.com/ogabrielrodrigues/imobiliary/internal/worker"
	"github.com/ogabrielrodrigues/imobiliary/util"
)

func main() {
	env := environment.LoadEnvironment()

	houses := reader.ReadCSV(env.TABLE_PATH)

	launch := launcher.New().Delete("--headless").MustLaunch()
	browser := rod.New().ControlURL(launch).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser)

	for i, house := range houses {
		houses[i].Debts = worker.Work(page, env.SAAEC_URL, house.ID)
	}

	go generator.Generate(houses)

	page.MustNavigate(fmt.Sprintf("http://%s/report", env.LOCAL_URL))
	page.MustWaitDOMStable()

	page.MustPDF(path.Join("reports", fmt.Sprintf("%s.pdf", util.CurrentDate())))
}
