package roomraccoon_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBookingNotesPost(t *testing.T) {
	req := client.NewBookingNotesPost()
	req.PathParams().BookingID = 1822
	body := req.NewRequestBody()
	body.Content = "Test from API"

	req.SetRequestBody(body)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
