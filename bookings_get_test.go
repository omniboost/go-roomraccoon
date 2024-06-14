package venuesuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-venuesuite"
)

func TestBookingsGet(t *testing.T) {
	req := client.NewBookingsGet()
	req.QueryParams().CreatedFrom = venuesuite.Date{time.Date(
		2024, 6, 14, 0, 0, 0, 0, time.UTC,
	)}
	req.QueryParams().CreatedUntil = venuesuite.Date{time.Date(
		2024, 6, 15, 0, 0, 0, 0, time.UTC,
	)}
	req.QueryParams().Status = []string{"request", "booked"}
	// req.QueryParams().Source = "3PY"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestBookingsGetAll(t *testing.T) {
	req := client.NewBookingsGet()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
