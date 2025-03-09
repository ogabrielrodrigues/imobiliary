package main

import (
	"fmt"
	"path/filepath"

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

	launch := launcher.New()
	if env.HEADLESS {
		launch.Delete("--headless")
	}

	if env.BROWSER_BIN != "" {
		launch.Bin(env.BROWSER_BIN)
	}

	browser := rod.New().ControlURL(launch.MustLaunch()).MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	fmt.Println("Initializing worker...")

	page := stealth.MustPage(browser)

	for i, house := range houses {
		fmt.Printf("Retrieving data from %s\n", house.Address)
		houses[i].Debts = worker.Work(page, env.SAAEC_URL, house.ID)
	}

	go generator.Generate(env.LOCAL_URL, houses)
	fmt.Println("Generating report...")

	page.MustNavigate(fmt.Sprintf("http://%s/report", env.LOCAL_URL))
	page.MustWaitDOMStable()

	report_path := filepath.Join("reports", fmt.Sprintf("%s.pdf", util.CurrentDate()))
	page.MustPDF(report_path)

	fmt.Printf("Report saved on: %s\n", report_path)
}
