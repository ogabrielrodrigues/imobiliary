package reader

import (
	"encoding/csv"
	"os"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
)

func ReadCSV(path string) []kind.House {
	file, err := os.Open(path)
	if err != nil {
		panic("err=" + err.Error())
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic("err=" + err.Error())
	}

	// Remove csv header.
	lines = lines[1:]

	var houses []kind.House
	for _, house := range lines {
		houses = append(houses, kind.House{
			Address: house[0],
			ID:      house[1],
		})
	}

	return houses
}
