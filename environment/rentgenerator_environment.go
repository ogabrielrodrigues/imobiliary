package environment

import (
	"flag"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared"
)

func LoadRentGeneratorEnvironment() *kind.RentGeneratorEnvironment {
	env := kind.RentGeneratorEnvironment{}

	rent_path := flag.String("path", "rentals.json", "path from rentals input data file.")
	port := flag.String("p", ":8080", "internal server port")

	flag.Parse()

	env.RENT_PATH = *rent_path
	env.LOCAL_URL = *port

	shared.Logln(shared.ColorGreen, "✓ Environment Variables sucessfully loaded!")

	return &env
}
