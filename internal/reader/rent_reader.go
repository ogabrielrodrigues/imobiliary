package reader

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared"
)

func ReadRentJSON(env *kind.RentGeneratorEnvironment) kind.RentTmplData {
	file, err := os.Open(env.RENT_PATH)
	if err != nil {
		shared.Logln(shared.ColorRed, "✗ ERROR opening JSON file.")
		os.Exit(1)
	}
	defer file.Close()

	var rent_data kind.RentTmplData
	err = json.NewDecoder(file).Decode(&rent_data)
	if err != nil {
		log.Fatal(err)
	}

	shared.Logln(shared.ColorGreen, "✓ JSON Rentals file read sucessfully!")

	return rent_data
}
