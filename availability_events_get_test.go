package venuesuite_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/omniboost/go-venuesuite"
)

func TestAvailabilityEventsGet(t *testing.T) {
	venueID, err := strconv.Atoi(os.Getenv("VENUE_ID"))
	if err != nil {
		t.Error(err)
	}

	req := client.NewAvailabilityEventsGet()
	req.QueryParams().Start = venuesuite.Date{time.Date(
		2024, 1, 1, 0, 0, 0, 0, time.UTC,
	)}
	req.QueryParams().End = venuesuite.Date{time.Date(
		2024, 5, 1, 0, 0, 0, 0, time.UTC,
	)}
	req.PathParams().VenueID = venueID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
