package reader

import (
	"encoding/csv"
	"os"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/util"
)

func ReadHouseCSV(path string) []kind.House {
	file, err := os.Open(path)
	if err != nil {
		util.Logln(util.ColorRed, "✗ ERROR opening csv file.")
		os.Exit(1)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		util.Logln(util.ColorRed, "✗ ERROR reading csv file.")
		os.Exit(1)
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

	util.Logln(util.ColorGreen, "✓ CSV file read sucessfully!")

	return houses
}
