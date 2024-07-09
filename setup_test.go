package roomraccoon_test

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"testing"

	roomraccoon "github.com/omniboost/go-roomraccoon"
)

var (
	client *roomraccoon.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("ROOMRACCOON_BASE_URL")
	apiKey := os.Getenv("ROOMRACCOON_API_KEY")
	hotelID, err := strconv.Atoi(os.Getenv("ROOMRACCOON_HOTEL_ID"))
	if err != nil {
		log.Fatal(err)
	}
	debug := os.Getenv("DEBUG")

	client = roomraccoon.NewClient(nil)
	client.SetApiKey(apiKey)
	client.SetHotelID(hotelID)
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
