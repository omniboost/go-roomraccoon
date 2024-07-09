package roomraccoon_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBookingActivitiesGet(t *testing.T) {
	req := client.NewBookingActivitiesGet()
	req.PathParams().BookingID = 1822
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestBookingActivitiesGetAll(t *testing.T) {
	req := client.NewBookingActivitiesGet()
	req.PathParams().BookingID = 1822
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
