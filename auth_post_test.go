package venuesuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAuthPost(t *testing.T) {
	req := client.NewAuthPost()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

