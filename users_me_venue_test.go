package venuesuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUsersMeVenue(t *testing.T) {
	req := client.NewUsersMeVenue()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}