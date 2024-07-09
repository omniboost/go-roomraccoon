package roomraccoon_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestVenueProductsGet(t *testing.T) {
	req := client.NewVenueProductsGet()
	venueID, err := strconv.Atoi(os.Getenv("VENUE_ID"))
	if err != nil {
		t.Error(err)
	}

	req.PathParams().VenueID = venueID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
