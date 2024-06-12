package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewVenueProducts() VenueProducts {
	r := VenueProducts{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VenueProducts struct {
	client      *Client
	queryParams *VenueProductsQueryParams
	pathParams  *VenueProductsPathParams
	method      string
	headers     http.Header
	requestBody VenueProductsBody
}

func (r VenueProducts) NewQueryParams() *VenueProductsQueryParams {
	return &VenueProductsQueryParams{}
}

type VenueProductsQueryParams struct {
}

func (p VenueProductsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VenueProducts) QueryParams() QueryParams {
	return r.queryParams
}

func (r VenueProducts) NewPathParams() *VenueProductsPathParams {
	return &VenueProductsPathParams{}
}

type VenueProductsPathParams struct {
	VenueID int `schema:"venue_id"`
}

func (p *VenueProductsPathParams) Params() map[string]string {
	return map[string]string{
		"venue_id": strconv.Itoa(p.VenueID),
	}
}

func (r *VenueProducts) PathParams() *VenueProductsPathParams {
	return r.pathParams
}

func (r *VenueProducts) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VenueProducts) SetMethod(method string) {
	r.method = method
}

func (r *VenueProducts) Method() string {
	return r.method
}

func (r VenueProducts) NewRequestBody() VenueProductsBody {
	return VenueProductsBody{}
}

type VenueProductsBody struct {
}

func (r *VenueProducts) RequestBody() *VenueProductsBody {
	return nil
}

func (r *VenueProducts) RequestBodyInterface() interface{} {
	return nil
}

func (r *VenueProducts) SetRequestBody(body VenueProductsBody) {
	r.requestBody = body
}

func (r *VenueProducts) NewResponseBody() *VenueProductsResponseBody {
	return &VenueProductsResponseBody{}
}

type VenueProductsResponseBody Products

func (r *VenueProducts) URL() *url.URL {
	u := r.client.GetEndpointURL("/venues/{{.venue_id}}/products", r.PathParams())
	return &u
}

func (r *VenueProducts) Do() (VenueProductsResponseBody, error) {
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
