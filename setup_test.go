package roomraccoon_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	roomraccoon "github.com/omniboost/go-roomraccoon"
)

var (
	client *roomraccoon.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	token := os.Getenv("VENUESUITE_TOKEN")
	debug := os.Getenv("DEBUG")

	client = roomraccoon.NewClient(nil)
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
