package roomraccoon_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBookingByIDGet(t *testing.T) {
	req := client.NewBookingByIDGet()
	req.PathParams().BookingID = 1805
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
