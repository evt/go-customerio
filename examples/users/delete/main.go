package main

import (
	"log"

	"github.com/evt/go-customerio/config"

	"github.com/evt/go-customerio"
)

func main() {
	cnf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	track := customerio.NewTrackClient(cnf.SiteID, cnf.ApiKey)

	err = track.Delete("customer_1")
	if err != nil {
		log.Fatal(err)
	}
}
