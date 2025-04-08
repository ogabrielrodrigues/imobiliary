package environment

import (
	"flag"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared"
)

func LoadReportEnvironment() *kind.ReportEnvironment {
	env := kind.ReportEnvironment{}

	headless := flag.Bool("headless", false, "enable browser headless mode.")
	path := flag.String("path", "houses.csv", "path from input table.")
	url := flag.String("url", "", "destination website url.")
	internal := flag.String("internal", "localhost:3000", "internal server address")
	bin := flag.String("bin", "", "bin path of user browser (optional).")
	out := flag.String("out", "reports", "reports output path. (optional).")

	flag.Parse()

	env.HEADLESS = *headless
	env.TABLE_PATH = *path
	env.SAAEC_URL = *url
	env.LOCAL_URL = *internal
	env.BROWSER_BIN = *bin
	env.REPORT_OUT = *out

	shared.Logln(shared.ColorGreen, "✓ Environment Variables sucessfully loaded!")

	return &env
}
