package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewVenueSpaces() VenueSpaces {
	r := VenueSpaces{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VenueSpaces struct {
	client      *Client
	queryParams *VenueSpacesQueryParams
	pathParams  *VenueSpacesPathParams
	method      string
	headers     http.Header
	requestBody VenueSpacesBody
}

func (r VenueSpaces) NewQueryParams() *VenueSpacesQueryParams {
	return &VenueSpacesQueryParams{}
}

type VenueSpacesQueryParams struct {
}

func (p VenueSpacesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VenueSpaces) QueryParams() QueryParams {
	return r.queryParams
}

func (r VenueSpaces) NewPathParams() *VenueSpacesPathParams {
	return &VenueSpacesPathParams{}
}

type VenueSpacesPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *VenueSpacesPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *VenueSpaces) PathParams() *VenueSpacesPathParams {
	return r.pathParams
}

func (r *VenueSpaces) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VenueSpaces) SetMethod(method string) {
	r.method = method
}

func (r *VenueSpaces) Method() string {
	return r.method
}

func (r VenueSpaces) NewRequestBody() VenueSpacesBody {
	return VenueSpacesBody{}
}

type VenueSpacesBody struct {
}

func (r *VenueSpaces) RequestBody() *VenueSpacesBody {
	return nil
}

func (r *VenueSpaces) RequestBodyInterface() interface{} {
	return nil
}

func (r *VenueSpaces) SetRequestBody(body VenueSpacesBody) {
	r.requestBody = body
}

func (r *VenueSpaces) NewResponseBody() *VenueSpacesResponseBody {
	return &VenueSpacesResponseBody{}
}

type VenueSpacesResponseBody Spaces

func (r *VenueSpaces) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/spaces", r.PathParams())
	return &u
}

func (r *VenueSpaces) Do() (VenueSpacesResponseBody, error) {
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
