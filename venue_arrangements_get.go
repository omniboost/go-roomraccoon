package roomraccoon

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-roomraccoon/utils"
)

func (c *Client) NewVenueArrangementsGet() VenueArrangementsGet {
	r := VenueArrangementsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VenueArrangementsGet struct {
	client      *Client
	queryParams *VenueArrangementsGetQueryParams
	pathParams  *VenueArrangementsGetPathParams
	method      string
	headers     http.Header
	requestBody VenueArrangementsGetBody
}

func (r VenueArrangementsGet) NewQueryParams() *VenueArrangementsGetQueryParams {
	return &VenueArrangementsGetQueryParams{}
}

type VenueArrangementsGetQueryParams struct {
}

func (p VenueArrangementsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VenueArrangementsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r VenueArrangementsGet) NewPathParams() *VenueArrangementsGetPathParams {
	return &VenueArrangementsGetPathParams{}
}

type VenueArrangementsGetPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *VenueArrangementsGetPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *VenueArrangementsGet) PathParams() *VenueArrangementsGetPathParams {
	return r.pathParams
}

func (r *VenueArrangementsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VenueArrangementsGet) SetMethod(method string) {
	r.method = method
}

func (r *VenueArrangementsGet) Method() string {
	return r.method
}

func (r VenueArrangementsGet) NewRequestBody() VenueArrangementsGetBody {
	return VenueArrangementsGetBody{}
}

type VenueArrangementsGetBody struct {
}

func (r *VenueArrangementsGet) RequestBody() *VenueArrangementsGetBody {
	return nil
}

func (r *VenueArrangementsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VenueArrangementsGet) SetRequestBody(body VenueArrangementsGetBody) {
	r.requestBody = body
}

func (r *VenueArrangementsGet) NewResponseBody() *VenueArrangementsGetResponseBody {
	return &VenueArrangementsGetResponseBody{}
}

type VenueArrangementsGetResponseBody Arrangements

func (r *VenueArrangementsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/arrangements", r.PathParams())
	return &u
}

func (r *VenueArrangementsGet) Do() (VenueArrangementsGetResponseBody, error) {
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
