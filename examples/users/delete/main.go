package main

import (
	"go-customerio/examples/users/config"
	"log"

	"github.com/customerio/go-customerio"
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
