package kind

import (
	"fmt"
	"os"
)

type Environment struct {
	HEADLESS    bool
	TABLE_PATH  string
	SAAEC_URL   string
	LOCAL_URL   string
	BROWSER_BIN string
}

func (e *Environment) ReadEnvironmentFromStdin() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("ERROR", "environment variables not provided")
		os.Exit(1)
	}

	args := os.Args[2:]

	if args[0] == "true" {
		e.HEADLESS = true
	} else {
		e.HEADLESS = false
	}

	e.TABLE_PATH = args[2]
	e.SAAEC_URL = args[4]
	e.LOCAL_URL = args[6]
	e.BROWSER_BIN = args[8]

	fmt.Println("Environment Variables sucessfully loaded!")
}
