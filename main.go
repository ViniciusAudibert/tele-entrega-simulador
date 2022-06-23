package main

import (
	"fmt"

	route "github.com/ViniciusAudibert/tele-entrega-simulador/application/route"
)

func main() {
	r := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	r.LoadPositions()
	stringJson, _ := r.ExportJsonPositons()

	fmt.Println(stringJson[1])
}
