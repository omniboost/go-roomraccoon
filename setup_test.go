package venuesuite_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	venuesuite "github.com/omniboost/go-venuesuite"
)

var (
	client *venuesuite.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	token := os.Getenv("VENUESUITE_TOKEN")
	debug := os.Getenv("DEBUG")

	client = venuesuite.NewClient(nil)
	client.SetToken(token)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
