package venuesuite

import (
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02 15:04:05"))
}

func (d *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02 15:04:05", value)
	if err == nil {
		return nil
	}

	// lastly try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	return err
}
