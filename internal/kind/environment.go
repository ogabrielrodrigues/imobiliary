package kind

import (
	"flag"

	"github.com/ogabrielrodrigues/imobiliary/util"
)

type Environment struct {
	HEADLESS    bool
	TABLE_PATH  string
	SAAEC_URL   string
	LOCAL_URL   string
	BROWSER_BIN string
	REPORT_OUT  string
}

func (e *Environment) ReadEnvironment() {
	headless := flag.Bool("headless", false, "enable browser headless mode.")
	path := flag.String("path", "houses.csv", "path from input table.")
	url := flag.String("url", "", "destination website url.")
	internal := flag.String("internal", "localhost:3000", "internal server address")
	bin := flag.String("bin", "", "bin path of user browser (optional).")
	out := flag.String("out", "reports", "reports output path. (optional).")

	flag.Parse()

	e.HEADLESS = *headless
	e.TABLE_PATH = *path
	e.SAAEC_URL = *url
	e.LOCAL_URL = *internal
	e.BROWSER_BIN = *bin
	e.REPORT_OUT = *out

	util.Logln(util.ColorGreen, "✓ Environment Variables sucessfully loaded!")
}
