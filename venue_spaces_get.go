package venuesuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewVenueSpacesGet() VenueSpacesGet {
	r := VenueSpacesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VenueSpacesGet struct {
	client      *Client
	queryParams *VenueSpacesGetQueryParams
	pathParams  *VenueSpacesGetPathParams
	method      string
	headers     http.Header
	requestBody VenueSpacesGetBody
}

func (r VenueSpacesGet) NewQueryParams() *VenueSpacesGetQueryParams {
	return &VenueSpacesGetQueryParams{}
}

type VenueSpacesGetQueryParams struct {
}

func (p VenueSpacesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VenueSpacesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r VenueSpacesGet) NewPathParams() *VenueSpacesGetPathParams {
	return &VenueSpacesGetPathParams{}
}

type VenueSpacesGetPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *VenueSpacesGetPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *VenueSpacesGet) PathParams() *VenueSpacesGetPathParams {
	return r.pathParams
}

func (r *VenueSpacesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VenueSpacesGet) SetMethod(method string) {
	r.method = method
}

func (r *VenueSpacesGet) Method() string {
	return r.method
}

func (r VenueSpacesGet) NewRequestBody() VenueSpacesGetBody {
	return VenueSpacesGetBody{}
}

type VenueSpacesGetBody struct {
}

func (r *VenueSpacesGet) RequestBody() *VenueSpacesGetBody {
	return nil
}

func (r *VenueSpacesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VenueSpacesGet) SetRequestBody(body VenueSpacesGetBody) {
	r.requestBody = body
}

func (r *VenueSpacesGet) NewResponseBody() *VenueSpacesGetResponseBody {
	return &VenueSpacesGetResponseBody{}
}

type VenueSpacesGetResponseBody Spaces

func (r *VenueSpacesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/spaces", r.PathParams())
	return &u
}

func (r *VenueSpacesGet) Do() (VenueSpacesGetResponseBody, error) {
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
