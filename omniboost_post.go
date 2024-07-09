package roomraccoon

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-roomraccoon/utils"
)

func (c *Client) NewOmniboostPost() OmniboostPost {
	r := OmniboostPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OmniboostPost struct {
	client      *Client
	queryParams *OmniboostPostQueryParams
	pathParams  *OmniboostPostPathParams
	method      string
	headers     http.Header
	requestBody OmniboostPostBody
}

func (r OmniboostPost) NewQueryParams() *OmniboostPostQueryParams {
	return &OmniboostPostQueryParams{}
}

type OmniboostPostQueryParams struct {
}

func (p OmniboostPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *OmniboostPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r OmniboostPost) NewPathParams() *OmniboostPostPathParams {
	return &OmniboostPostPathParams{}
}

type OmniboostPostPathParams struct {
}

func (p *OmniboostPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OmniboostPost) PathParams() *OmniboostPostPathParams {
	return r.pathParams
}

func (r *OmniboostPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OmniboostPost) SetMethod(method string) {
	r.method = method
}

func (r *OmniboostPost) Method() string {
	return r.method
}

func (r OmniboostPost) NewRequestBody() OmniboostPostBody {
	return OmniboostPostBody{}
}

type OmniboostPostBody struct {
	HotelID   int    `json:"hotelId"`
	ApiKey    string `json:"apiKey"`
	SDateFrom Date   `json:"sDateFrom"`
	SDateTo   Date   `json:"sDateTo"`
}

func (r *OmniboostPost) RequestBody() *OmniboostPostBody {
	return &r.requestBody
}

func (r *OmniboostPost) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *OmniboostPost) SetRequestBody(body OmniboostPostBody) {
	r.requestBody = body
}

func (r *OmniboostPost) NewResponseBody() *OmniboostPostResponseBody {
	return &OmniboostPostResponseBody{}
}

type OmniboostPostResponseBody struct {
	Success bool `json:"success"`
	Data    []struct {
		Transactions []struct {
			NetAmount   StringFloat `json:"netAmount"`
			TaxAmount   StringFloat `json:"taxAmount"`
			GrossAmount StringFloat `json:"grossAmount"`
			Description string      `json:"description"`
			Category    string      `json:"category"`
			Ledger      struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"ledger"`
			Taxes []struct {
				Name       string      `json:"name"`
				Percentage StringFloat `json:"percentage"`
				Amount     StringFloat `json:"amount"`
			} `json:"taxes"`

			TaxCode string `json:"taxCode"`
			Room    struct {
				Name     string `json:"name"`
				Number   string `json:"number"`
				Category string `json:"category"`
			} `json:"room"`

			Date DateTime `json:"date"`
		} `json:"transactions"`

		Payments []struct {
			Amount            StringFloat `json:"amount"`
			Date              DateTime    `json:"date"`
			Description       string      `json:"description"`
			Initials          string      `json:"initials"`
			InvoiceNumber     string      `json:"invoiceNumber"`
			Method            string      `json:"method"`
			ReservationNumber string      `json:"reservationNumber"`
			Ledger            struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"ledger"`
		} `json:"payments"`
		SalesEntry struct {
			InvoiceNumber string   `json:"invoiceNumber"`
			InvoiceDate   DateTime `json:"invoiceDate"`
			InvoiceTotal  float64  `json:"invoiceTotal"`
			InvoiceType   string   `json:"invoiceType"`
			PaymentTotal  float64  `json:"paymentTotal"`
		} `json:"salesEntry"`
		Reservation struct {
			StartDate         Date   `json:"startDate"`
			EndDate           Date   `json:"endDate"`
			ReservationNumber string `json:"reservationNumber"`
		} `json:"reservation"`
		Debtor struct {
			Name            string `json:"name"`
			Email           string `json:"email"`
			UniqueID        string `json:"uniqueId"`
			Gender          string `json:"gender"`
			TelephoneNumber string `json:"telephoneNumber"`
			Ascription      string `json:"ascription"`
		} `json:"debtor,omitempty"`
		DebtorCode string `json:"debtorCode"`
	} `json:"data"`
}

func (r *OmniboostPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/api/omniboost", r.PathParams())
	return &u
}

func (r *OmniboostPost) Do() (OmniboostPostResponseBody, error) {
	r.requestBody.ApiKey = r.client.apiKey
	r.requestBody.HotelID = r.client.hotelID

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
