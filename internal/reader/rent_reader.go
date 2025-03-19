package reader

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/util"
)

func ReadRentJSON(env *kind.RentGeneratorEnvironment) kind.RentTmplData {
	file, err := os.Open(env.RENT_PATH)
	if err != nil {
		util.Logln(util.ColorRed, "✗ ERROR opening JSON file.")
		os.Exit(1)
	}
	defer file.Close()

	var rent_data kind.RentTmplData
	err = json.NewDecoder(file).Decode(&rent_data)
	if err != nil {
		log.Fatal(err)
	}

	util.Logln(util.ColorGreen, "✓ JSON Rentals file read sucessfully!")

	return rent_data
}
