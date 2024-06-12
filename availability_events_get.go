package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewAvailabilityEventsGet() AvailabilityEventsGet {
	r := AvailabilityEventsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AvailabilityEventsGet struct {
	client      *Client
	queryParams *AvailabilityEventsGetQueryParams
	pathParams  *AvailabilityEventsGetPathParams
	method      string
	headers     http.Header
	requestBody AvailabilityEventsGetBody
}

func (r AvailabilityEventsGet) NewQueryParams() *AvailabilityEventsGetQueryParams {
	return &AvailabilityEventsGetQueryParams{}
}

type AvailabilityEventsGetQueryParams struct {
	Start Date `schema:"start"`
	End   Date `schema:"end"`
}

func (p AvailabilityEventsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AvailabilityEventsGet) QueryParams() *AvailabilityEventsGetQueryParams {
	return r.queryParams
}

func (r AvailabilityEventsGet) NewPathParams() *AvailabilityEventsGetPathParams {
	return &AvailabilityEventsGetPathParams{}
}

type AvailabilityEventsGetPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *AvailabilityEventsGetPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *AvailabilityEventsGet) PathParams() *AvailabilityEventsGetPathParams {
	return r.pathParams
}

func (r *AvailabilityEventsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AvailabilityEventsGet) SetMethod(method string) {
	r.method = method
}

func (r *AvailabilityEventsGet) Method() string {
	return r.method
}

func (r AvailabilityEventsGet) NewRequestBody() AvailabilityEventsGetBody {
	return AvailabilityEventsGetBody{}
}

type AvailabilityEventsGetBody struct {
}

func (r *AvailabilityEventsGet) RequestBody() *AvailabilityEventsGetBody {
	return nil
}

func (r *AvailabilityEventsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *AvailabilityEventsGet) SetRequestBody(body AvailabilityEventsGetBody) {
	r.requestBody = body
}

func (r *AvailabilityEventsGet) NewResponseBody() *AvailabilityEventsGetResponseBody {
	return &AvailabilityEventsGetResponseBody{}
}

type AvailabilityEventsGetResponseBody struct {
	Data AvailabilityEvents `json:"data"`
}

func (r *AvailabilityEventsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/availability_events", r.PathParams())
	return &u
}

func (r *AvailabilityEventsGet) Do() (AvailabilityEventsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	token, err := r.client.Token()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
