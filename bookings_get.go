package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewBookingsGet() BookingsGet {
	r := BookingsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BookingsGet struct {
	client      *Client
	queryParams *BookingsGetQueryParams
	pathParams  *BookingsGetPathParams
	method      string
	headers     http.Header
	requestBody BookingsGetBody
}

func (r BookingsGet) NewQueryParams() *BookingsGetQueryParams {
	return &BookingsGetQueryParams{}
}

type BookingsGetQueryParams struct {
}

func (p BookingsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BookingsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r BookingsGet) NewPathParams() *BookingsGetPathParams {
	return &BookingsGetPathParams{}
}

type BookingsGetPathParams struct{}

func (p *BookingsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BookingsGet) PathParams() *BookingsGetPathParams {
	return r.pathParams
}

func (r *BookingsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BookingsGet) SetMethod(method string) {
	r.method = method
}

func (r *BookingsGet) Method() string {
	return r.method
}

func (r BookingsGet) NewRequestBody() BookingsGetBody {
	return BookingsGetBody{}
}

type BookingsGetBody struct {
}

func (r *BookingsGet) RequestBody() *BookingsGetBody {
	return nil
}

func (r *BookingsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *BookingsGet) SetRequestBody(body BookingsGetBody) {
	r.requestBody = body
}

func (r *BookingsGet) NewResponseBody() *BookingsGetResponseBody {
	return &BookingsGetResponseBody{}
}

type BookingsGetResponseBody struct {
	Data Bookings `json:"data"`
	Meta struct {
		Pagination Pagination `json:"pagination"`
	} `json:"meta"`
}

func (r *BookingsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/bookings", r.PathParams())
	return &u
}

func (r *BookingsGet) Do() (BookingsGetResponseBody, error) {
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