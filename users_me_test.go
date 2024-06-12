package venuesuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUsersMe(t *testing.T) {
	req := client.NewUsersMe()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

