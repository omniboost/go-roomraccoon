package venuesuite

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Venue     int       `json:"venue"`
	LastLogin time.Time `json:"last_login"`
	CreatedOn time.Time `json:"created_on"`
	Detail    struct {
		FirstName      string `json:"first_name"`
		Preposition    string `json:"preposition"`
		LastName       string `json:"last_name"`
		Avatar         string `json:"avatar"`
		Phone          string `json:"phone"`
		EmailSignature string `json:"email_signature"`
		Lang           string `json:"lang"`
		Locale         string `json:"locale"`
	} `json:"detail"`
}
