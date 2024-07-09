package roomraccoon

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-roomraccoon/utils"
)

func (c *Client) NewAvailabilityCalendarsGet() AvailabilityCalendarsGet {
	r := AvailabilityCalendarsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AvailabilityCalendarsGet struct {
	client      *Client
	queryParams *AvailabilityCalendarsGetQueryParams
	pathParams  *AvailabilityCalendarsGetPathParams
	method      string
	headers     http.Header
	requestBody AvailabilityCalendarsGetBody
}

func (r AvailabilityCalendarsGet) NewQueryParams() *AvailabilityCalendarsGetQueryParams {
	return &AvailabilityCalendarsGetQueryParams{}
}

type AvailabilityCalendarsGetQueryParams struct {
}

func (p AvailabilityCalendarsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AvailabilityCalendarsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r AvailabilityCalendarsGet) NewPathParams() *AvailabilityCalendarsGetPathParams {
	return &AvailabilityCalendarsGetPathParams{}
}

type AvailabilityCalendarsGetPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *AvailabilityCalendarsGetPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *AvailabilityCalendarsGet) PathParams() *AvailabilityCalendarsGetPathParams {
	return r.pathParams
}

func (r *AvailabilityCalendarsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AvailabilityCalendarsGet) SetMethod(method string) {
	r.method = method
}

func (r *AvailabilityCalendarsGet) Method() string {
	return r.method
}

func (r AvailabilityCalendarsGet) NewRequestBody() AvailabilityCalendarsGetBody {
	return AvailabilityCalendarsGetBody{}
}

type AvailabilityCalendarsGetBody struct {
}

func (r *AvailabilityCalendarsGet) RequestBody() *AvailabilityCalendarsGetBody {
	return nil
}

func (r *AvailabilityCalendarsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *AvailabilityCalendarsGet) SetRequestBody(body AvailabilityCalendarsGetBody) {
	r.requestBody = body
}

func (r *AvailabilityCalendarsGet) NewResponseBody() *AvailabilityCalendarsGetResponseBody {
	return &AvailabilityCalendarsGetResponseBody{}
}

type AvailabilityCalendarsGetResponseBody struct {
	Data []any `json:"data"`
}

func (r *AvailabilityCalendarsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/availability_calendars", r.PathParams())
	return &u
}

func (r *AvailabilityCalendarsGet) Do() (AvailabilityCalendarsGetResponseBody, error) {
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
