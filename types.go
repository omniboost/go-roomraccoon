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

type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Venues []Venue

type Product struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Component   string   `json:"component"`
	Description string   `json:"description,omitempty"`
	Activated   string   `json:"activated,omitempty"`
	Images      []string `json:"images"`
	Pricing     struct {
		Unit          string `json:"unit"`
		Included      int    `json:"included"`
		Excluded      int    `json:"excluded"`
		TaxPercentage int    `json:"tax_percentage"`
	} `json:"pricing"`
	Quantity  int    `json:"quantity,omitempty"`
	Type      string `json:"type"`
	Gradation any    `json:"gradation,omitempty"`
}

type Products []Product
