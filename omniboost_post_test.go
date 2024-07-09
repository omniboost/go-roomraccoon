package roomraccoon_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-roomraccoon"
)

func TestOmniboostPost(t *testing.T) {
	req := client.NewOmniboostPost()
	start := time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2024, 7, 9, 0, 0, 0, 0, time.Local)
	for start.Before(end) {
		req.RequestBody().SDateFrom = roomraccoon.Date{start}
		req.RequestBody().SDateTo = roomraccoon.Date{start.AddDate(0, 0, 1)}

		resp, err := req.Do()
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(b))
		start = start.AddDate(0, 0, 1)
	}

}
