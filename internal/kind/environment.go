package kind

type ReportEnvironment struct {
	HEADLESS    bool
	TABLE_PATH  string
	SAAEC_URL   string
	LOCAL_URL   string
	BROWSER_BIN string
	REPORT_OUT  string
}

type ServerEnvironment struct {
	SERVER_ADDR   string
	DATABASE_HOST string
	DATABASE_PORT string
	DATABASE_NAME string
	DATABASE_USER string
	DATABASE_PWD  string
}
