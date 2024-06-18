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

func TestVenueBookingsGet(t *testing.T) {
	req := client.NewVenueBookingsGet()
	venueID, err := strconv.Atoi(os.Getenv("VENUE_ID"))
	if err != nil {
		t.Error(err)
	}

	req.PathParams().VenueID = venueID
	req.QueryParams().CreatedFrom = venuesuite.Date{time.Date(
		2024, 1, 1, 0, 0, 0, 0, time.UTC,
	)}
	req.QueryParams().CreatedUntil = venuesuite.Date{time.Date(
		2024, 5, 1, 0, 0, 0, 0, time.UTC,
	)}
	req.QueryParams().Status = []string{"draft", "booked"}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestVenueBookingsGetAll(t *testing.T) {
	req := client.NewVenueBookingsGet()
	venueID, err := strconv.Atoi(os.Getenv("VENUE_ID"))
	if err != nil {
		t.Error(err)
	}

	req.PathParams().VenueID = venueID
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}