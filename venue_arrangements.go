package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewVenueArrangements() VenueArrangements {
	r := VenueArrangements{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VenueArrangements struct {
	client      *Client
	queryParams *VenueArrangementsQueryParams
	pathParams  *VenueArrangementsPathParams
	method      string
	headers     http.Header
	requestBody VenueArrangementsBody
}

func (r VenueArrangements) NewQueryParams() *VenueArrangementsQueryParams {
	return &VenueArrangementsQueryParams{}
}

type VenueArrangementsQueryParams struct {
}

func (p VenueArrangementsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VenueArrangements) QueryParams() QueryParams {
	return r.queryParams
}

func (r VenueArrangements) NewPathParams() *VenueArrangementsPathParams {
	return &VenueArrangementsPathParams{}
}

type VenueArrangementsPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *VenueArrangementsPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *VenueArrangements) PathParams() *VenueArrangementsPathParams {
	return r.pathParams
}

func (r *VenueArrangements) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VenueArrangements) SetMethod(method string) {
	r.method = method
}

func (r *VenueArrangements) Method() string {
	return r.method
}

func (r VenueArrangements) NewRequestBody() VenueArrangementsBody {
	return VenueArrangementsBody{}
}

type VenueArrangementsBody struct {
}

func (r *VenueArrangements) RequestBody() *VenueArrangementsBody {
	return nil
}

func (r *VenueArrangements) RequestBodyInterface() interface{} {
	return nil
}

func (r *VenueArrangements) SetRequestBody(body VenueArrangementsBody) {
	r.requestBody = body
}

func (r *VenueArrangements) NewResponseBody() *VenueArrangementsResponseBody {
	return &VenueArrangementsResponseBody{}
}

type VenueArrangementsResponseBody Arrangements

func (r *VenueArrangements) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/arrangements", r.PathParams())
	return &u
}

func (r *VenueArrangements) Do() (VenueArrangementsResponseBody, error) {
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
