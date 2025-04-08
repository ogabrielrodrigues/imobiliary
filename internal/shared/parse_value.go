package shared

import (
	"strconv"
	"strings"
)

func ParseValue(value string) float32 {
	value = strings.ReplaceAll(value, ",", ".")

	parsed, _ := strconv.ParseFloat(value, 32)

	return float32(parsed)
}
