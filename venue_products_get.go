package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewVenueProductsGet() VenueProductsGet {
	r := VenueProductsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VenueProductsGet struct {
	client      *Client
	queryParams *VenueProductsGetQueryParams
	pathParams  *VenueProductsGetPathParams
	method      string
	headers     http.Header
	requestBody VenueProductsGetBody
}

func (r VenueProductsGet) NewQueryParams() *VenueProductsGetQueryParams {
	return &VenueProductsGetQueryParams{}
}

type VenueProductsGetQueryParams struct {
}

func (p VenueProductsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VenueProductsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r VenueProductsGet) NewPathParams() *VenueProductsGetPathParams {
	return &VenueProductsGetPathParams{}
}

type VenueProductsGetPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *VenueProductsGetPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *VenueProductsGet) PathParams() *VenueProductsGetPathParams {
	return r.pathParams
}

func (r *VenueProductsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VenueProductsGet) SetMethod(method string) {
	r.method = method
}

func (r *VenueProductsGet) Method() string {
	return r.method
}

func (r VenueProductsGet) NewRequestBody() VenueProductsGetBody {
	return VenueProductsGetBody{}
}

type VenueProductsGetBody struct {
}

func (r *VenueProductsGet) RequestBody() *VenueProductsGetBody {
	return nil
}

func (r *VenueProductsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VenueProductsGet) SetRequestBody(body VenueProductsGetBody) {
	r.requestBody = body
}

func (r *VenueProductsGet) NewResponseBody() *VenueProductsGetResponseBody {
	return &VenueProductsGetResponseBody{}
}

type VenueProductsGetResponseBody Products

func (r *VenueProductsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/products", r.PathParams())
	return &u
}

func (r *VenueProductsGet) Do() (VenueProductsGetResponseBody, error) {
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
