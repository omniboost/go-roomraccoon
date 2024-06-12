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
	req.QueryParams().Start  = venuesuite.Date{time.Date(
		2024, 1, 1, 0, 0, 0, 0, time.UTC,
	)}
	req.QueryParams().End  = venuesuite.Date{time.Date(
		2024, 5, 1, 0, 0, 0, 0, time.UTC,
	)}
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
