package main

import (
	"flag"
	"fmt"
	"github.com/joesweeny/statshub/internal/config"
	"github.com/joesweeny/statshub/internal/bootstrap"
	"os"
)

const country = "country"

var option = flag.String("option", "", "Provide the model name to process")

func main() {
	app := bootstrap.Bootstrap{Config: config.GetConfig()}

	flag.Parse()

	switch *option {
	case country:
		app.GetCountryService().Process()
		fmt.Println("Countries proceeded successfully")
		os.Exit(0)
	default:
		fmt.Println("The option provided is not supported")
		os.Exit(1)
	}
}

