package venuesuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBookingsGet(t *testing.T) {
	req := client.NewBookingsGet()
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
