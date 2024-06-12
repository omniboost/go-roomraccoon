package venuesuite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-venuesuite/utils"
)

func (c *Client) NewUsersMeVenue() UsersMeVenue {
	r := UsersMeVenue{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type UsersMeVenue struct {
	client      *Client
	queryParams *UsersMeVenueQueryParams
	pathParams  *UsersMeVenuePathParams
	method      string
	headers     http.Header
	requestBody UsersMeVenueBody
}

func (r UsersMeVenue) NewQueryParams() *UsersMeVenueQueryParams {
	return &UsersMeVenueQueryParams{}
}

type UsersMeVenueQueryParams struct {
}

func (p UsersMeVenueQueryParams) ToURLValues() (url.Values, error) {
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

func (r *UsersMeVenue) QueryParams() QueryParams {
	return r.queryParams
}

func (r UsersMeVenue) NewPathParams() *UsersMeVenuePathParams {
	return &UsersMeVenuePathParams{}
}

type UsersMeVenuePathParams struct {
}

func (p *UsersMeVenuePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *UsersMeVenue) PathParams() *UsersMeVenuePathParams {
	return r.pathParams
}

func (r *UsersMeVenue) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *UsersMeVenue) SetMethod(method string) {
	r.method = method
}

func (r *UsersMeVenue) Method() string {
	return r.method
}

func (r UsersMeVenue) NewRequestBody() UsersMeVenueBody {
	return UsersMeVenueBody{}
}

type UsersMeVenueBody struct {
}

func (r *UsersMeVenue) RequestBody() *UsersMeVenueBody {
	return nil
}

func (r *UsersMeVenue) RequestBodyInterface() interface{} {
	return nil
}

func (r *UsersMeVenue) SetRequestBody(body UsersMeVenueBody) {
	r.requestBody = body
}

func (r *UsersMeVenue) NewResponseBody() *UsersMeVenueResponseBody {
	return &UsersMeVenueResponseBody{}
}

type UsersMeVenueResponseBody Venues

func (r *UsersMeVenue) URL() *url.URL {
	u := r.client.GetEndpointURL("/users/me/venues", r.PathParams())
	return &u
}

func (r *UsersMeVenue) Do() (UsersMeVenueResponseBody, error) {
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
