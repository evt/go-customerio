package main

import (
	"go-customerio/examples/users/config"
	"log"
	"time"

	"github.com/customerio/go-customerio"
)

func main() {
	cnf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	track := customerio.NewTrackClient(cnf.SiteID, cnf.ApiKey)

	err = track.Identify("customer_1", map[string]interface{}{
		"email":      "bob@example.com",
		"created_at": time.Now().Unix(),
		"first_name": "Bob",
		"plan":       "basic",
	})
	if err != nil {
		log.Fatal(err)
	}
}
