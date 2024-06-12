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

type Arrangement struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Type          string   `json:"type"`
	Description   string   `json:"description,omitempty"`
	Category      string   `json:"category"`
	Component     string   `json:"component"`
	Images        []string `json:"images,omitempty"`
	Weight        int      `json:"weight"`
	Duration      int      `json:"duration"`
	Spaces        []string `json:"spaces"`
	MinimumGuests int      `json:"minimum_guests"`
	MaximumGuests int      `json:"maximum_guests"`
	Products      struct {
		Catering []struct {
			ID          int      `json:"id"`
			Title       string   `json:"title"`
			Category    string   `json:"category"`
			Component   string   `json:"component"`
			Gradation   any      `json:"gradation"`
			Description string   `json:"description,omitempty"`
			Images      []string `json:"images,omitempty"`
			Pricing     struct {
				Unit          string `json:"unit"`
				Included      int    `json:"included"`
				Excluded      int    `json:"excluded"`
				TaxPercentage int    `json:"tax_percentage"`
			} `json:"pricing"`
			Activated any    `json:"activated"`
			Type      string `json:"type"`
			Quantity  int    `json:"quantity"`
		} `json:"catering"`
		Equipment []any `json:"equipment"`
	} `json:"products"`
	Pricing []struct {
		Excluded      int    `json:"excluded"`
		Included      int    `json:"included"`
		TaxPercentage int    `json:"tax_percentage"`
		Unit          string `json:"unit"`
	} `json:"pricing"`
	Discounts           any  `json:"discounts"`
	SpacePriceIncluded  bool `json:"space_price_included"`
	SpaceAlwaysIncluded int  `json:"space_always_included"`
	Activated           int  `json:"activated"`
}

type Arrangements []Arrangement
