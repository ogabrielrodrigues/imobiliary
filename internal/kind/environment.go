package kind

import (
	"os"
)

type Environment struct {
	TABLE_PATH string
	SAAEC_URL  string
	LOCAL_URL  string
}

func (e *Environment) ReadEnvironmentFromStdin() {
	args := os.Args[2:]

	if len(os.Args[1:]) != 0 {
		e.TABLE_PATH = args[0]
		e.SAAEC_URL = args[2]
		e.LOCAL_URL = args[4]
	}
}
