package main

import (
	"fmt"
	"path/filepath"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/reader"
	"github.com/ogabrielrodrigues/imobiliary/internal/report/generator"
	"github.com/ogabrielrodrigues/imobiliary/internal/report/worker"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared"
)

func main() {
	env := environment.LoadReportEnvironment()

	houses := reader.ReadHouseCSV(env.TABLE_PATH)

	launch := launcher.New()
	if !env.HEADLESS {
		launch.Delete("--headless")
	}

	if env.BROWSER_BIN != "" {
		launch.Bin(env.BROWSER_BIN)
	}

	browser := rod.New().ControlURL(launch.MustLaunch()).MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	shared.Logln(shared.ColorGreen, "✓ Initializing worker...")

	page := stealth.MustPage(browser)

	for i, house := range houses {
		shared.Log(shared.ColorDefault, fmt.Sprintf("Retrieving data from: %s. ", house.Address))
		houses[i].Debts = worker.Work(page, env.SAAEC_URL, house.ID)
	}

	go generator.Generate(env.LOCAL_URL, houses)
	shared.Logln(shared.ColorDefault, "\nGenerating report...")

	page.MustNavigate(fmt.Sprintf("http://%s/report", env.LOCAL_URL))
	page.MustWaitDOMStable()

	report_path := filepath.Join(env.REPORT_OUT, fmt.Sprintf("%s.pdf", shared.CurrentDate()))
	page.MustPDF(report_path)

	shared.Logln(shared.ColorGreen, fmt.Sprintf("Report saved on: %s\n", report_path))
}
