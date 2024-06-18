package venuesuite

import (
	"encoding/json"
	"strings"
	"time"
)

type CommaSeparatedQueryParam []string

func (t CommaSeparatedQueryParam) MarshalSchema() string {
	return strings.Join(t, ",")
}

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
	Pricing     Pricing  `json:"pricing"`
	Quantity    int      `json:"quantity,omitempty"`
	Type        string   `json:"type"`
	Gradation   any      `json:"gradation,omitempty"`
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
			Pricing     Pricing  `json:"pricing"`
			Activated   any      `json:"activated"`
			Type        string   `json:"type"`
			Quantity    int      `json:"quantity"`
		} `json:"catering"`
		Equipment []any `json:"equipment"`
	} `json:"products"`
	Pricing             Pricings `json:"pricing"`
	Discounts           any      `json:"discounts"`
	SpacePriceIncluded  bool     `json:"space_price_included"`
	SpaceAlwaysIncluded int      `json:"space_always_included"`
	Activated           int      `json:"activated"`
}

type Arrangements []Arrangement

type Space struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Component   string `json:"component"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	EventTypes  string `json:"event_types"`
	Pricing     struct {
		Hour struct {
			Excluded float64 `json:"excluded"`
			Included float64 `json:"included"`
		} `json:"hour"`
		Day struct {
			Excluded float64 `json:"excluded"`
			Included float64 `json:"included"`
		} `json:"day"`
		Daypart struct {
			Excluded float64 `json:"excluded"`
			Included float64 `json:"included"`
		} `json:"daypart"`
		TaxPercentage string `json:"tax_percentage"`
		Type          string `json:"type"`
		Min           int    `json:"min"`
	} `json:"pricing"`
	Images     []string `json:"images"`
	Facilities []any    `json:"facilities"`
	Capacity   struct {
		Max int `json:"max"`
		Min int `json:"min"`
	} `json:"capacity"`
	Equipment []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Category    any    `json:"category"`
		Component   any    `json:"component"`
		Description any    `json:"description"`
		Activated   any    `json:"activated"`
	} `json:"equipment"`
	Venue struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"venue"`
	Setups []struct {
		Label string `json:"label"`
		Min   int    `json:"min"`
		Max   int    `json:"max"`
	} `json:"setups"`
	LinkedSpaces []any `json:"linked_spaces"`
}

type Spaces []Space

type AvailabilityEvents []AvailabilityEvent

type AvailabilityEvent struct {
	End              DateTime    `json:"end"`
	ID               int         `json:"id"`
	Source           string      `json:"source"`
	SourceIdentifier interface{} `json:"source_identifier"`
	Space            struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Title string `json:"title"`
	} `json:"space"`
	Start DateTime `json:"start"`
	Title string   `json:"title"`
}

type Pagination struct {
	Total       int      `json:"total"`
	Count       int      `json:"count"`
	PerPage     int      `json:"per_page"`
	CurrentPage int      `json:"current_page"`
	TotalPages  int      `json:"total_pages"`
	Links       struct{} `json:"links"`
}

type Pricings []Pricing

type Pricing struct {
	Unit          string  `json:"unit"`
	Included      float64 `json:"included"`
	Excluded      float64 `json:"excluded"`
	TaxPercentage float64 `json:"tax_percentage"`
}

type BookingSlots []BookingSlot

type BookingSlot struct {
	ID        int       `json:"id"`
	BookingID int       `json:"booking_id"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Guests    int       `json:"guests"`
	Packages  []struct {
		ID            int      `json:"id"`
		Title         string   `json:"title"`
		Type          string   `json:"type"`
		Description   any      `json:"description"`
		Category      string   `json:"category"`
		Component     string   `json:"component"`
		Images        any      `json:"images"`
		Weight        int      `json:"weight"`
		Duration      int      `json:"duration"`
		Spaces        []string `json:"spaces"`
		MinimumGuests int      `json:"minimum_guests"`
		MaximumGuests int      `json:"maximum_guests"`
		Products      []struct {
			ID          int     `json:"id"`
			Title       string  `json:"title"`
			Category    string  `json:"category"`
			Component   string  `json:"component"`
			Gradation   any     `json:"gradation"`
			Description any     `json:"description"`
			Images      any     `json:"images"`
			Pricing     Pricing `json:"pricing"`
			Activated   any     `json:"activated"`
			Type        string  `json:"type"`
			Quantity    int     `json:"quantity"`
		} `json:"products"`
		Pricing             Pricings `json:"pricing"`
		Discounts           any      `json:"discounts"`
		SpacePriceIncluded  bool     `json:"space_price_included"`
		SpaceAlwaysIncluded int      `json:"space_always_included"`
		Activated           int      `json:"activated"`
		Quote               struct {
			Guests              int        `json:"guests"`
			BasePrice           int        `json:"base_price"`
			Price               float64    `json:"price"`
			Discount            float64    `json:"discount"`
			Tax                 BookingTax `json:"tax"`
			TotalPrice          int        `json:"total_price"`
			BasePricePerPerson  int        `json:"base_price_per_person"`
			TotalPricePerPerson int        `json:"total_price_per_person"`
		} `json:"quote"`
		Quantity   int       `json:"quantity"`
		ExternalID IntString `json:"external_id"`
	} `json:"packages"`
	Catering []struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Category    string    `json:"category"`
		Component   string    `json:"component"`
		Gradation   any       `json:"gradation"`
		Description string    `json:"description"`
		Images      []string  `json:"images"`
		Pricing     Pricing   `json:"pricing"`
		Activated   any       `json:"activated"`
		Type        string    `json:"type"`
		Quantity    int       `json:"quantity"`
		ExternalID  IntString `json:"external_id"`
	} `json:"catering"`
	Equipment []struct {
		ID          int      `json:"id"`
		Title       string   `json:"title"`
		Category    string   `json:"category"`
		Component   string   `json:"component"`
		Description any      `json:"description"`
		Activated   any      `json:"activated"`
		Images      []string `json:"images"`
		Pricing     Pricing  `json:"pricing"`
		Quantity    int      `json:"quantity"`
		Type        string   `json:"type"`
	} `json:"equipment"`
	Extras []any `json:"extras"`
	Spaces []struct {
		ID          int    `json:"id"`
		Category    string `json:"category"`
		Component   string `json:"component"`
		Title       string `json:"title"`
		Description any    `json:"description"`
		EventTypes  string `json:"event_types"`
		Pricing     struct {
			// @TODO: add extra unmarshal method
			// Hour struct {
			// 	Excluded float64 `json:"excluded"`
			// 	Included float64 `json:"included"`
			// } `json:"hour"`
			// Day struct {
			// 	Excluded float64 `json:"excluded"`
			// 	Included float64 `json:"included"`
			// } `json:"day"`
			// Daypart struct {
			// 	Excluded float64 `json:"excluded"`
			// 	Included float64 `json:"included"`
			// } `json:"daypart"`
			TaxPercentage StringFloat `json:"tax_percentage"`
			Type          string      `json:"type"`
			Min           int         `json:"min"`
		} `json:"pricing"`
		Images     []string `json:"images"`
		Facilities []any    `json:"facilities"`
		Capacity   struct {
			Max float64 `json:"max"`
			Min float64 `json:"min"`
		} `json:"capacity"`
		Equipment []struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Category    any    `json:"category"`
			Component   any    `json:"component"`
			Description any    `json:"description"`
			Activated   any    `json:"activated"`
		} `json:"equipment"`
		Venue struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"venue"`
		Setups []struct {
			Label string `json:"label"`
			Min   int    `json:"min"`
			Max   int    `json:"max"`
		} `json:"setups"`
		LinkedSpaces []any     `json:"linked_spaces"`
		ExternalID   IntString `json:"external_id"`
	} `json:"spaces"`
	Program []any `json:"program"`
}

type BookingTax map[string]float64

func (d *BookingTax) UnmarshalJSON(text []byte) error {
	// Check if it is not an object, because empty summaries are arrays instead of objects...
	if string(text)[0] != '{' {
		return nil
	}

	type Alias BookingTax
	alias := Alias(*d)
	err := json.Unmarshal(text, &alias)
	if err != nil {
		return err
	}

	return nil
}

type BookingInvoiceSlotSummary struct {
	Guests              int        `json:"guests"`
	BasePrice           float64    `json:"base_price"`
	Price               float64    `json:"price"`
	Discount            float64    `json:"discount"`
	Tax                 BookingTax `json:"tax"`
	TotalPrice          float64    `json:"total_price"`
	BasePricePerPerson  float64    `json:"base_price_per_person"`
	TotalPricePerPerson float64    `json:"total_price_per_person"`
}

func (d *BookingInvoiceSlotSummary) UnmarshalJSON(text []byte) error {
	// Check if it is not an object, because empty summaries are arrays instead of objects...
	if string(text)[0] != '{' {
		return nil
	}

	type Alias BookingInvoiceSlotSummary
	alias := Alias(*d)
	err := json.Unmarshal(text, &alias)
	if err != nil {
		return err
	}

	return nil
}

type BookingInvoiceSlots []BookingInvoiceSlot

type BookingInvoiceSlot struct {
	SlotID int       `json:"slot_id"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Guests int       `json:"guests"`
	Items  []struct {
		Category       string  `json:"category"`
		Product        int     `json:"product"`
		Title          string  `json:"title"`
		Description    string  `json:"description"`
		Quantity       int     `json:"quantity"`
		UnitNetPrice   float64 `json:"unit_net_price"`
		UnitGrossPrice float64 `json:"unit_gross_price"`
		GrossPrice     float64 `json:"gross_price"`
		NetPrice       float64 `json:"net_price"`
		TaxRate        float64 `json:"tax_rate"`
		TaxPercentage  float64 `json:"tax_percentage"`
		DiscountNet    float64 `json:"discount_net"`
		DiscountBy     string  `json:"discount_by"`
	} `json:"items"`
	Summary BookingInvoiceSlotSummary `json:"summary"`
}

type BookingInvoice struct {
	Summary struct {
		Guests              int        `json:"guests"`
		BasePrice           int        `json:"base_price"`
		Price               int        `json:"price"`
		Discount            int        `json:"discount"`
		Tax                 BookingTax `json:"tax"`
		TotalPrice          int        `json:"total_price"`
		BasePricePerPerson  int        `json:"base_price_per_person"`
		TotalPricePerPerson int        `json:"total_price_per_person"`
	} `json:"summary"`
	Slots BookingInvoiceSlots `json:"slots"`
}

type Booking struct {
	ID            int    `json:"id"`
	Version       int    `json:"version"`
	GroupID       int    `json:"group_id"`
	Reference     string `json:"reference"`
	PurchaseOrder string `json:"purchase_order,omitempty"`
	CostCenter    string `json:"cost_center,omitempty"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Status        string `json:"status"`
	Setup         string `json:"setup"`
	Guests        int    `json:"guests"`
	Period        struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"period"`
	Discount            []any          `json:"discount"`
	Slots               []BookingSlot  `json:"slots"`
	Notes               string         `json:"notes,omitempty"`
	Program             []any          `json:"program,omitempty"`
	Currency            string         `json:"currency"`
	Invoice             BookingInvoice `json:"invoice"`
	CreatedOn           time.Time      `json:"created_on"`
	UpdatedOn           time.Time      `json:"updated_on"`
	OrganizationContact struct {
		ID           int    `json:"id"`
		Email        string `json:"email"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Phone        string `json:"phone"`
		Lang         string `json:"lang"`
		Role         string `json:"role"`
		Organization struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"organization"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"organization_contact"`
	CancelDate time.Time `json:"cancel_date"`
	Venue      struct {
		ID       int      `json:"id"`
		Name     string   `json:"name"`
		Terms    string   `json:"terms"`
		Features []string `json:"features"`
	} `json:"venue"`
	Organization struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Street      string `json:"street"`
		StreetNr    string `json:"street_nr"`
		StreetNrAdd any    `json:"street_nr_add"`
		Region      string `json:"region"`
		Province    string `json:"province"`
		City        string `json:"city"`
		Country     string `json:"country"`
	} `json:"organization"`
	VenueContact struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Detail   struct {
			FirstName   string `json:"first_name"`
			Preposition string `json:"preposition"`
			LastName    string `json:"last_name"`
			Avatar      string `json:"avatar"`
			Phone       string `json:"phone,omitempty"`
		} `json:"detail"`
	} `json:"venue_contact"`
}

type Bookings []Booking
