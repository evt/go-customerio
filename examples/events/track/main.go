package main

import (
	"github.com/evt/go-customerio/config"
	"log"

	"github.com/evt/go-customerio"
)

func main() {
	cnf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	track := customerio.NewTrackClient(cnf.SiteID, cnf.ApiKey)

	err = track.Track("5", "purchase", map[string]interface{}{
		"type": "socks",
		"price": "14.99",
	})
	if err != nil {
		log.Fatal(err)
	}
}
