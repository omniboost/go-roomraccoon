package venuesuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewVenueBookingsGet() VenueBookingsGet {
	r := VenueBookingsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VenueBookingsGet struct {
	client      *Client
	queryParams *VenueBookingsGetQueryParams
	pathParams  *VenueBookingsGetPathParams
	method      string
	headers     http.Header
	requestBody VenueBookingsGetBody
}

func (r VenueBookingsGet) NewQueryParams() *VenueBookingsGetQueryParams {
	return &VenueBookingsGetQueryParams{}
}

type VenueBookingsGetQueryParams struct {
	Source       CommaSeparatedQueryParam `schema:"source,omitempty"`
	Status       CommaSeparatedQueryParam `schema:"status,omitempty"`
	PeriodStart  Date                     `schema:"period[start],omitempty"`
	PeriodEnd    Date                     `schema:"period[end],omitempty"`
	CreatedFrom  Date                     `schema:"created_from,omitempty"`
	CreatedUntil Date                     `schema:"created_until,omitempty"`
	Limit        int                      `schema:"limit,omitempty"`
	Page         int                      `schema:"page,omitempty"`
}

func (p VenueBookingsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *VenueBookingsGet) QueryParams() *VenueBookingsGetQueryParams {
	return r.queryParams
}

func (r VenueBookingsGet) NewPathParams() *VenueBookingsGetPathParams {
	return &VenueBookingsGetPathParams{}
}

type VenueBookingsGetPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *VenueBookingsGetPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *VenueBookingsGet) PathParams() *VenueBookingsGetPathParams {
	return r.pathParams
}

func (r *VenueBookingsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VenueBookingsGet) SetMethod(method string) {
	r.method = method
}

func (r *VenueBookingsGet) Method() string {
	return r.method
}

func (r VenueBookingsGet) NewRequestBody() VenueBookingsGetBody {
	return VenueBookingsGetBody{}
}

type VenueBookingsGetBody struct {
}

func (r *VenueBookingsGet) RequestBody() *VenueBookingsGetBody {
	return nil
}

func (r *VenueBookingsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VenueBookingsGet) SetRequestBody(body VenueBookingsGetBody) {
	r.requestBody = body
}

func (r *VenueBookingsGet) NewResponseBody() *VenueBookingsGetResponseBody {
	return &VenueBookingsGetResponseBody{}
}

type VenueBookingsGetResponseBody struct {
	Data Bookings `json:"data"`
	Meta struct {
		Pagination Pagination `json:"pagination"`
	} `json:"meta"`
}

func (r *VenueBookingsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/bookings", r.PathParams())
	return &u
}

func (r *VenueBookingsGet) Do() (VenueBookingsGetResponseBody, error) {
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

func (r *VenueBookingsGet) All() (Bookings, error) {
	bookings := Bookings{}
	response := *r.NewResponseBody()
	for {
		resp, err := r.Do()
		if err != nil {
			return bookings, err
		}

		// Break out of loop when no bookings are found
		if len(resp.Data) == 0 {
			break
		}

		// Add bookings to list
		bookings = append(bookings, resp.Data...)

		if response.Meta.Pagination.CurrentPage == response.Meta.Pagination.TotalPages {
			break
		}

		// Increment page number
		r.QueryParams().Page = response.Meta.Pagination.CurrentPage + 1
	}

	return bookings, nil
}
