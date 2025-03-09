package reader

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
)

func ReadCSV(path string) []kind.House {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR", "error opening csv file.")
		os.Exit(1)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("ERROR", "error reading csv file.")
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

	fmt.Println("CSV file read sucessfully!")
	return houses
}
