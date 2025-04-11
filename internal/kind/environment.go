package kind

type ReportEnvironment struct {
	HEADLESS    bool
	TABLE_PATH  string
	SAAEC_URL   string
	LOCAL_URL   string
	BROWSER_BIN string
	REPORT_OUT  string
}

type APIEnvironment struct {
	SERVER_ADDR   string
	DATABASE_HOST string
	DATABASE_PORT string
	DATABASE_NAME string
	DATABASE_USER string
	DATABASE_PWD  string
	SECRET_KEY    string
	CORS_ORIGIN   string
}

type RentGeneratorEnvironment struct {
	RENT_PATH string
	LOCAL_URL string
}
