package util

import (
	"fmt"
	"time"
)

func CurrentDate() string {
	t := time.Now()
	return fmt.Sprintf("%d-%d-%d",
		t.Day(),
		t.Month(),
		t.Year(),
	)
}
