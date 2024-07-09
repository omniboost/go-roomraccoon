package roomraccoon_test

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/omniboost/go-roomraccoon"
)

func TestAvailabilityEventPost(t *testing.T) {
	venueID, err := strconv.Atoi(os.Getenv("VENUE_ID"))
	if err != nil {
		t.Error(err)
	}

	req := client.NewAvailabilityEventPost()
	req.PathParams().VenueID = venueID
	req.RequestBody().Start = roomraccoon.DateTime{time.Date(
		2024, 6, 12, 14, 21, 0, 0, time.UTC,
	)}
	req.RequestBody().End = roomraccoon.DateTime{time.Date(
		2024, 6, 12, 18, 0, 0, 0, time.UTC,
	)}
	req.RequestBody().Title = "Test event"
	req.RequestBody().SpaceID = 4795
	// req.RequestBody().SpaceID = 12
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
